
let tabNav = new Vue({
    el: '#tabNav',
    data: {
        act: "encrypt"
    },
    computed: {
        active: {
            // getter
            get: function () {
                return this.act
            },
            // setter
            set: function (newValue) {
                this.act = newValue
                switch(this.act) {
                    case "decrypt":
                        decryptForm.show = true
                        encryptForm.show = false
                        infoContainer.show = false
                        break;
                    case "info":
                        decryptForm.show = false
                        encryptForm.show = false
                        infoContainer.show = true
                        break;
                    default:
                        decryptForm.show = false
                        infoContainer.show = false
                        encryptForm.show = true
                }
            }

        },
        show:function () {
            return this.error.length > 0
        }
    }
})