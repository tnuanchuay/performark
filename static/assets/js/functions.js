
/* Background Images
-------------------------------------------------------------------*/
var  sectionImage = jQuery('#section').data('background-image');
var  resultImage = jQuery('#result').data('background-image');
if (sectionImage) {  jQuery('#section').css({ 'background-image':'url(' + sectionImage + ')' }); };
if (resultImage) {  jQuery('#result').css({ 'background-image':'url(' + resultImage + ')' }); };

/* Background Images End
-------------------------------------------------------------------*/



/* Document Ready function
-------------------------------------------------------------------*/
jQuery(document).ready(function($) {

	"use strict";


    /* Window Height Resize
    -------------------------------------------------------------------*/
    var windowheight = jQuery(window).height();
    if(windowheight > 650)
    {
         $('.pattern').removeClass('height-resize');
    }
    /* Window Height Resize End
    -------------------------------------------------------------------*/



	/* Main Menu
	-------------------------------------------------------------------*/
	$('#main-menu #headernavigation').onePageNav({
		currentClass: 'active',
		changeHash: false,
		scrollSpeed: 750,
		scrollThreshold: 0.5,
		scrollOffset: 0,
		filter: '',
		easing: 'swing'
	});

	/* Main Menu End
	-------------------------------------------------------------------*/




});

/* Document Ready function End
-------------------------------------------------------------------*/


/* Preloder
-------------------------------------------------------------------*/
$(window).load(function () {
    "use strict";
    $("#loader").fadeOut();
    $("#preloader").delay(350).fadeOut("slow");
});
 /* Preloder End
-------------------------------------------------------------------*/
