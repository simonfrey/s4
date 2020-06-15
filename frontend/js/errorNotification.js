
let errorNotification = new Vue({
    el: '#errorNotification',
    data: {
        err: ""
    },
    computed: {
        error: {
            // getter
            get: function () {
                return this.err
            },
            // setter
            set: function (newValue) {
                this.err = newValue
            }

        },
        show:function () {
            return this.error.length > 0
        }
    }
})