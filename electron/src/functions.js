export default {
    isGoodDate: function(dt){
        let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
        return reGoodDate.test(dt);
    }
}