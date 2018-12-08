
// vue component for publish
var config = new Vue({
    el: '#publish-changes',
    data: {
        commitMessage:'',
        notification: '',
        isPublishInProgress: false
    },
    methods: {
        publishChanges: function () {
            if (this.commitMessage.length < 1) {
                alert("Please enter a commit message")
                return
            }
            this.isPublishInProgress = true
            this.notification = "Publishing the changes. Pls wait ..."

            axios
                .get('api/v1/publish?commitMessage=' + this.commitMessage)
                .then(response => {
                    this.notification = "Successfully published on to github @" + getDateAndTime()
                    this.isPublishInProgress = false
                })
                .catch(error => {
                    this.notification = "Failed to publish @" + getDateAndTime()
                    this.isPublishInProgress = false
                })

        }
    }

})
