var optScales = {
yAxes: [{
    ticks: {beginAtZero:false, callback:function(label, index, labels){
        if(label/1000000 > 1){
            return label / 1000000 + 'M'
        }else if(label/1000 > 1){
            return label / 1000 + 'k'
        }else {
            return label
        }
    }},
    }]
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
};

var optTitle = {
display: true,
	 text:"Request / Second",
	 position:"top",
	 fontColor: '#fff',
	 fontSize:16,
};