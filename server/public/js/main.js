// hide main view
$("#main").hide()
$(document).ready(function(){

    $("#main").show()
    $(".menu-item").click(function(){
        $(".menu-item").css("background-color", "white")
        $(".menu-item").css("color", "royalblue")
        
        var selected = "."+this.classList[1]
        
        $("li"+selected).css("background-color", "royalblue")
        $("li"+selected).css("color", "white")
        // change main view
        changeView(selected)
    })
})

// Change the main-view for selected menu-item
function changeView(selectedClass) {
    $(".main-view").hide()

    if (selectedClass == ".split") {
        // Add class names 
        $("div.code").addClass("split-left")
        $("div.preview").addClass("split-right")
        
        // Display code & preview pages
        $("div.code").show()
        $("div.preview").show()

        return
    } else {
        $("div.code").removeClass("split-left")
        $("div.preview").removeClass("split-right")
    }

	$("div"+selectedClass).show()
    // $("div.preview").show()
}

function getDateAndTime() {
    return new Date()
}