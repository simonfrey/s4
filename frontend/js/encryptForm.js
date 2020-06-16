let encryptForm = new Vue({
    el: '#encryptForm',
    data: {
        iString: "",
        n:2,
        k:2,
        oArray:Array("",""),
        aes:true,
        iShow: true,
    },
    computed: {
        outArray: {
            // getter
            get: function () {
                return this.oArray
            },
            // setter
            set: function (newValue) {
                this.oArray = newValue
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
        useAES:{

            // getter
            get: function () {
                return Boolean(this.aes)
            },
            // setter
            set: function (newValue) {
                this.aes = Boolean(newValue)
                this.distribute()
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
                if (this.n < this.requiredShareCount){
                    this.requiredShareCount = this.n
                }

                this.distribute()
            }

        },
        requiredShareCount: {
            // getter
            get: function () {
                if (this.k > this.share){
                    return Number(this.shares)
                }
                return Number(this.k)
            },
            // setter
            set: function (newValue) {
                if (newValue > this.shares){
                    return
                }

                this.k = Number(newValue)
                this.distribute()
            }

        },
        inputString: {
            // getter
            get: function () {
                return String(this.iString)
            },
            // setter
            set: function (newValue) {
                this.iString = String(newValue)
                this.distribute()
            }

        },
    },
    methods:{
        distribute:function () {
            if (this.inputString.length === 0){
                this.outArray = Array(this.shares).fill("")
                return
            }
            if (this.shares <= 0 || this.requiredShareCount<=0||this.requiredShareCount>this.shares){
                return
            }

            res = Distribute_fours((String(btoa(this.inputString))),Number(this.shares),Number(this.requiredShareCount),Boolean(this.useAES))
            if (typeof res === 'string'){
                errorNotification.error = res
                this.outArray = Array(this.shares).fill("")
                return
            }
            this.outArray = res
            errorNotification.error = ""
        }
    }
})
