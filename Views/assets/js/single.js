var app = new Vue({
    delimiters: ['$((', '))'],
    el: "#app",
    data: {
        baseUrl: document.querySelector("input[name=baseUrl]").value,
        linkId: document.querySelector("input[name=link_id]").value,
        modal: false,
        dates: [],
        is_edit: false,

        feeds: []

    },

    mounted: function() {
        console.log('ok');
        console.log(this.baseUrl);
        this.getDates();
    },

    methods: {

        getDates: function() {

            axios.get(`${this.baseUrl}/api/links/${this.linkId}/dates`).then(function(result) {

                if (result.data.error) {
                    alert(result.data.error)
                    return false
                }

                //console.log(result.data)

                app.dates = result.data.data;

                if (app.dates) {
                    app.getFeeds(app.dates[0].date)
                }
            });
        },

        getFeeds: function (date) {
            axios.get(`${this.baseUrl }/api/links/${this.linkId}/dates/${date}/feeds`).then(function(result) {

                if (result.data.error) {
                    alert(result.data.error)
                    return false
                }

                //console.log(result.data)

                app.feeds = result.data.data;

            });
        }

    }

});