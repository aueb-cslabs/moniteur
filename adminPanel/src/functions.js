export default {
    isGoodDate: function(dt){
        let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
        return reGoodDate.test(dt);
    },

    isValidTime: function(timeStr) {
        return (timeStr.search(/^\d{2}:\d{2}$/) !== -1) &&
        (timeStr.substr(0,2) >= 0 && timeStr.substr(0,2) <= 24) &&
        (timeStr.substr(3,2) >= 0 && timeStr.substr(3,2) <= 59);
    },

    showSidebar: function () {
        if (!window.sideOpen){
            let mw = window.matchMedia( "(max-width: 500px) ");
            if (mw.matches){
                document.getElementById("sidebar").style.width = "100%";
                document.getElementById("main").style.marginLeft = "100%";
            }
            else {
                document.getElementById("sidebar").style.width = "280px";
                document.getElementById("main").style.marginLeft = "280px";
            }
            window.sideOpen = true;
        }
        else {
            document.getElementById("sidebar").style.width = "0px";
            document.getElementById("main").style.marginLeft = "0px";
            window.sideOpen = false;
        }
    }
}