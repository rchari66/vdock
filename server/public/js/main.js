
$(document).ready(function(){
    $(".menu-item").click(function(){
        $(".menu-item").css("background-color", "white");
        $(".menu-item").css("color", "royalblue");
        
        var selected = "."+this.classList[1]
        
        $("li"+selected).css("background-color", "royalblue");
        $("li"+selected).css("color", "white");
        // change main view
        changeView(selected);
    });
});

// Change the main-view for selected menu-item
function changeView(selectedClass) {
	$(".main-view").hide();
	$("div"+selectedClass).show();
    // $("div.preview").show();
}
