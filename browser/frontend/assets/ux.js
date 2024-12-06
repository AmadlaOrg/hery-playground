$(document).ready(function() {
    // This initializes the tabs
    $("#tabs").tabs({
        activate: function(event, ui) {
            // Remove 'pf-m-current' from all PatternFly tabs and panels
            $(".pf-v6-c-tabs__item").removeClass("pf-m-current");
            //$(".pf-v6-c-tabs__panel").removeClass("pf-m-current");

            // Add 'pf-m-current' to the selected tab and its associated panel
            ui.newTab.addClass("pf-m-current");
            ui.newPanel.addClass("pf-m-current");
        }
    });
});
