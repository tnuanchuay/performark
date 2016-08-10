var rps = $("#rps");
var tps = $("#tps");
var latency = $("#latency");
var thread = $("#thread");
var requests = $("#requests");
var transfer = $("#transfer");
var errorc1 = $("#errorc1");
var errorc10 = $("#errorc10");
var errorc100 = $("#errorc100");
var errorc1k = $("#errorc1k");
var errorc10k = $("#errorc10k");
var errorc100k = $("#errorc100k");
var rpe = $("#rpe");
var rpe2 = $("#rpe2");

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

var transfers = new Chart(transfer, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Transfer',
data: {{.tt}},
backgroundColor: 'rgba(255, 159, 64, 0.2)',
borderColor: 'rgba(255, 159, 64, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"Transfer",
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

var ec = {{.ec}};
var er = {{.er}};
var ew = {{.ew}};
var et = {{.et}};
var ex = {{.ex}};
var ee = [];
var maxError = {
    i : 0,
    num : 0,
}
for(var i = 0 ; i < ec.length ; i++){
    ee[i] = [];
    ee[i].push(ec[i]);
    ee[i].push(er[i]);
    ee[i].push(ew[i]);
    ee[i].push(et[i]);
    ee[i].push(ex[i]);

    var err = ec[i]+er[i]+ew[i]+et[i]+ex[i]

    if(maxError.num <= err){
        maxError.i = i
        maxError.num = err
    }
}
//
//console.log(ec, er, ew, et, ex, error)

var errorc1c = new Chart(errorc1, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[0],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C1 Request And Error",
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
var errorc10c = new Chart(errorc10, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[1],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C10 Request And Error",
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
var errorc100c = new Chart(errorc100, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[2],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C100 Request And Error",
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
var errorc1kc = new Chart(errorc1k, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[3],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C1K Request And Error",
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
var errorc10kc = new Chart(errorc10k, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[4],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C10K Request And Error",
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
var errorc100kc = new Chart(errorc100k, {
type: 'pie',
data: {
labels: ["Socket Error Connect", "Socket Error Read", "Socket Error Write", "Socket Error Timeout", "Socket Error 2xx or 3xx"],
datasets: [{
label: 'Request',
data: ee[5],
backgroundColor: [
                 'rgba(255, 99, 132, 0.2)',
                 'rgba(54, 162, 235, 0.2)',
                 'rgba(255, 206, 86, 0.2)',
                 'rgba(255, 159, 64, 0.2)',
                 'rgba(75, 192, 192, 0.2)'
                 ],
borderColor: [
             'rgba(255, 99, 132, 0.2)',
             'rgba(54, 162, 235, 0.2)',
             'rgba(255, 206, 86, 0.2)',
             'rgba(255, 159, 64, 0.2)',
             'rgba(75, 192, 192, 0.2)'
             ],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"C100K Request And Error",
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

var rpec = new Chart(rpe, {
type: 'line',
data: {
labels: xLabel,
datasets: [{
label: 'Request',
data: {{.r}},
backgroundColor: 'rgba(153, 102, 255, 0.2)',
borderColor: 'rgba(153, 102, 255, 1)',
borderWidth: 4
},
{
label: 'Socket Error',
data: {{.e}},
backgroundColor: 'rgba(75, 192, 192, 0.2)',
borderColor: 'rgba(75, 192, 192, 1)',
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:"Request And Error",
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

var data = {{.r}}[maxError.i];
var selectedError = "";
switch(maxError.i){
    case 0 :
        selectedError = "C1"
        break;
    case 1 :
        selectedError = "C10"
        break;
    case 2 :
        selectedError = "C100"
        break;
    case 3 :
        selectedError = "C1k"
        break;
    case 4 :
        selectedError = "C10k"
        break;
    case 5 :
        selectedError = "C100k"
        break;
}
var dataSet = [data, maxError.num];
var rpe2c = new Chart(rpe2, {
type: 'pie',
data: {
labels: ["Success Request", "Error"],
datasets: [{
data: dataSet,
backgroundColor: ['rgba(153, 102, 255, 0.2)','rgba(75, 192, 192, 0.2)'],
borderColor: ['rgba(153, 102, 255, 1)','rgba(75, 192, 192, 1)'],
borderWidth: 4
}]
},
options: {
responsive: true,
	    title:{
display: true,
	 text:selectedError + " Request And Error Ratio",
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

