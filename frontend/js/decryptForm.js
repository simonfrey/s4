const base64regex = /^([0-9a-zA-Z+/]{4})*(([0-9a-zA-Z+/]{2}==)|([0-9a-zA-Z+/]{3}=))?$/;

let decryptForm = new Vue({
    el: '#decryptForm',
    data: {
        oString: "",
        n:2,
        oArray:Array("",""),
        iShow:false,
    },
    computed: {
        inFields: {
            // getter
            get: function () {
                return this.oArray
            },
            // setter
            set: function (newValue) {
                this.oArray = newValue
                this.recover()
            }

        },
        show: {
            // getter
            get: function () {
                return this.iShow
            },
            // setter
            set: function (newValue) {
                this.iShow = newValue
            }

        },
        shares: {
            // getter
            get: function () {
                return Number(this.n)
            },
            // setter
            set: function (newValue) {
                this.n = Number(newValue)
                this.inFields = Array(this.n).fill("")
                this.outputString = ""
                this.recover()
            }

        },
        outputString: {
            // getter
            get: function () {
                return String(this.oString)
            },
            // setter
            set: function (newValue) {
                this.oString = String(newValue)
            }

        },
    },
    methods:{
        recover:function () {
            console.log("recover")

           res = Recover_fours(this.inFields)

            if (!base64regex.test(res)){
                this.outputString = ""
                errorNotification.error = res
                return;
            }

            this.outputString = atob(res)
            errorNotification.error = ""


        },
        setInField: function (index,event) {
            newA = this.inFields
            newA[index] = String(event.target.value)
            this.inFields = newA
        }
    }
})