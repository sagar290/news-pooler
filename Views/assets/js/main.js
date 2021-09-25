var app = new Vue({
    delimiters: ['$((', '))'],
    el: "#app",
    data: {
        baseUrl: document.querySelector("input[name=baseUrl]").value,
        modal: false,
        body: {
            title: '',
            url: '',
            url_selector: '',
            description_selector: '',
        },

        urls: []

    },

    mounted: function() {
        console.log('ok');
        console.log(this.baseUrl);
        this.getUlrs()
    },

    methods: {

        addUrl: function() {

            if (!app.body.title ||
                !app.body.url ||
                !app.body.url_selector ||
                !app.body.description_selector ||
                !this.isValidURL(app.body.url)
            ) {
                alert('cant add link')

                return false
            }

            axios.post(this.baseUrl + '/api/links', app.body).then(function(result) {
                if (result.data.error) {
                    // app.errorMsg = score.data.message;
                } else {
                    data = result.data;
                    console.log(result.data);
                }
            });
        },

        getUlrs: function() {
            axios.get(this.baseUrl + '/api/links').then(function(result) {
                if (result.data.error) {
                    // app.errorMsg = score.data.message;
                } else {
                    app.urls = result.data.data;
                    console.log(result.data);
                }
            });
        },

        isValidURL: function(string) {
            var res = string.match(/(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g);
            return (res !== null)
        }
    }

});