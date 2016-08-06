var rps = $("#rps");
var tps = $("#tps");
var latency = $("#latency");
var thread = $("#thread");
var requests = $("#requests");

var global = Chart.defaults.global;
global.defaultFontColor = '#FFF';
var xLabel = ["c1", "c10", "c100", "c1k", "c10k", "c100k"]
var optTitle = {
display: true,
	 text:"Request / Second",
	 position:"top",
	 fontColor: '#fff',
	 fontSize:16,
};
var optScales = {
yAxes: [{
ticks: {
beginAtZero:false
       }
       }],
};

var optLegend = {
position:"bottom",
	 labels:{
fontSize:12,
	 fontColor:"#fff",
	 },
};

var optHover = {
mode: 'dataset'
};

var optTooltips = {
mode: 'label',
}

var rpsc = new Chart(rps, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Req/Sec',
data: {{.rps}},
backgroundColor: 'rgba(255, 99, 132, 0.2)',
borderColor: 'rgba(255,99,132,1)',
borderWidth: 4
}]
},
options: {
responsive: true,
title:{
display: true,
text:"Request / Second",
position:"top",
fontColor: '#fff',
fontSize:16,
},
legend: optLegend,
	scales: optScales,
	hover: optHover,
	tooltips: optTooltips
	}
});

var tpsc = new Chart(tps, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Transfer/Sec',
data: {{.tps}},
backgroundColor: 'rgba(54, 162, 235, 0.2)',
borderColor:'rgba(54, 162, 235, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
title:{
display: true,
text:"Transfer / Second",
position:"top",
fontColor: '#fff',
fontSize:16,
},
legend: optLegend,
	scales: optScales,
	hover: optHover,
	tooltips: optTooltips
	}
});

var latencyc = new Chart(latency, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Latency Max',
data: {{.lm}},
backgroundColor: 'rgba(255, 206, 86, 0.2)',
borderColor: 'rgba(255, 206, 86, 1)',
borderWidth: 4
},{
label: 'Latency Avg',
data: {{.la}},
backgroundColor: 'rgba(75, 192, 192, 0.2)',
borderColor: 'rgba(75, 192, 192, 1)',
borderWidth: 4
},{
label: 'Latency Stdev',
data: {{.ls}},
backgroundColor: 'rgba(153, 102, 255, 0.2)',
borderColor: 'rgba(153, 102, 255, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"Latency",
	 position:"top",
	 fontColor: '#fff',
	 fontSize:16,
	    },
legend: optLegend,
	scales: optScales,
	hover: optHover,
	tooltips: optTooltips
	 }
});

var threadc = new Chart(thread, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Max Request/Sec',
data: {{.tm}},
backgroundColor: 'rgba(255, 206, 86, 0.2)',
borderColor: 'rgba(255, 206, 86, 1)',
borderWidth: 4
},{
label: 'Avg Request/Sec',
data: {{.ta}},
backgroundColor: 'rgba(75, 192, 192, 0.2)',
borderColor: 'rgba(75, 192, 192, 1)',
borderWidth: 4
},{
label: 'Stdev Latency',
data: {{.ts}},
backgroundColor: 'rgba(153, 102, 255, 0.2)',
borderColor: 'rgba(153, 102, 255, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"One Thread Stat",
	 position:"top",
	 fontColor: '#fff',
	 fontSize:16,
	    },
legend: optLegend,
	scales: optScales,
	hover: optHover,
	tooltips: optTooltips
	 }
});

var requestsc = new Chart(requests, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Requests',
data: {{.r}},
backgroundColor: 'rgba(255, 159, 64, 0.2)',
borderColor: 'rgba(255, 159, 64, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"Requests",
	 position:"top",
	 fontColor: '#fff',
	 fontSize:16,
	    },
legend: optLegend,
	scales: optScales,
	hover: optHover,
	tooltips: optTooltips
	 }
});

