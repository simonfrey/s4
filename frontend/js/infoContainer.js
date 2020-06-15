
let infoContainer = new Vue({
    el: '#infoContainer',
    data: {
        iShow:false,
    },
    computed: {
        show: {
            // getter
            get: function () {
                return this.iShow
            },
            // setter
            set: function (newValue) {
                this.iShow = newValue
            }

        }
    }
})