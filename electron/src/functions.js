export default {
    isGoodDate: function(dt){
        let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
        return reGoodDate.test(dt);
    },

    isValidTime: function(timeStr) {
        return (timeStr.search(/^\d{2}:\d{2}$/) !== -1) &&
        (timeStr.substr(0,2) >= 0 && timeStr.substr(0,2) <= 24) &&
        (timeStr.substr(3,2) >= 0 && timeStr.substr(3,2) <= 59);
    }
}