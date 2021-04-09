var loadChart = function(chartElementId, formElementId, dataUrl, options) {
	var chart = null;

	var formElement = document.getElementById(formElementId);
	formElement.onsubmit = function() {
		var minValue = parseFloat(formElement.querySelector('#minValue').value);
		var dimension = formElement.querySelector("#dimension").value;
		buildLineChart({ minValue: minValue, dimension: dimension });
		return false;
	}

	var buildLineChart = function(extraOptions) {
		var options = Object.assign(extraOptions, options);
		var url = new URL(dataUrl);
		var urlParams = new URLSearchParams(url.search.slice(1));
		var startDate = new Date();
		startDate.setHours(startDate.getHours() - 6);
		urlParams.append('start', startDate.toISOString());
		var endDate = new Date();
		urlParams.append('end', endDate.toISOString());
		urlParams.append('dimension', options.dimension);
	
		var horizontalLabels = [];
		var datasetData = [];
		var datasetLabel = options.datasetLabel ? options.datasetLabel : "Value"
		var minValue = options.minValue ? options.minValue : 0;
		var maxValue = minValue + 10;
		
		getJSON(dataUrl + '?' + urlParams, function(err, json) {
			if (err) {
				alert(err);
				return;
			}
			
			for (var i = 0; i < json.data.length; i++) {
				var date = new Date(json.data[i].timestamp);
				var value = json.data[i].data;
				horizontalLabels.push(date.toLocaleString());
				datasetData.push(value);
				if (value > maxValue) {
					maxValue = value;
				}
			}
	
			if (chart) {
				chart.destroy();
			}
			chart = baseLineChart(chartElementId, datasetData, datasetLabel, horizontalLabels, minValue, maxValue);
		});
	}

	var minValue = parseFloat(formElement.querySelector('#minValue').value);
	var dimension = formElement.querySelector("#dimension").value;
	buildLineChart({ minValue: minValue, dimension: dimension });
}

var baseLineChart = function(elementId, data, datasetLabel, horizontalLabels, verticalScaleMin, verticalScaleMax) {
	var ctx = document.getElementById(elementId).getContext('2d');	
	return new Chart(ctx, {
		type: 'line',
		data: {
			labels: horizontalLabels,
			datasets: [{
				label: datasetLabel,
				data: data,
				backgroundColor: ['rgba(255, 99, 132, 0.2)'],
				borderColor: ['rgba(255, 99, 132, 1)'],
				borderWidth: 1
			}]
		},
		options: {
			scales: {
				y: {
					min: verticalScaleMin,
					max: verticalScaleMax
				}
			}
		}
	});
}

var getJSON = function(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function() {
      var status = xhr.status;
      if (status === 200) {
        callback(null, xhr.response);
      } else {
        callback(status, xhr.response);
      }
    };
    xhr.send();
};