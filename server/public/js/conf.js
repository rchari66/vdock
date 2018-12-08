
// vue component for config
var config = new Vue({
    el: '#set-config',
    data: {
        blogRepo: '',
        siteRepo: '',
        userId: '',
        emailId: '',
        password: '',
        repos: [],
        notification: '',
        isFetching: false,
        isUpdateInProgress: false

    },
    computed: {

    },
    methods: {
        fetchRepos: function () {
            var name = this.userId;
            if (name.length < 1) {
                alert("Please enter userId")
                return
            }
            // Disable fetch button
            this.isFetching = true
            this.notification = 'Fetching repository list ...' + " @" + getDateAndTime()
            axios
                .get('api/v1/conf/getRepos?userId=' + name)
                .then(response => {
                    // empty list
                    this.repos.splice(0, this.repos.length);

                    data = response.data
                    for (i = 0; i < data.length; i++) {
                        // this.$set(this.repos, i, data[i])
                        this.repos.push(data[i])
                    }
                    this.notification = 'Fetched repository list!!' + " @" + getDateAndTime()
                    this.isFetching = false
                })
                .catch(error => {
                    alert("Either invalid git userid or empty repository list")
                    this.isFetching = false
                    Console.log("Failed to get repository list")
                })
        },
        updateConfig: function () {
            // validate input length & content
            if (this.userId.length < 0 || this.password.length < 0 || this.emailId.length < 0) {
                alert("Please enter valid user details")
                return
            }
            if (this.blogRepo.length < 0 || this.blogRepo.length < 0) {
                alert("Please enter valid repository details")
                return
            }
            // Disable buttons 
            this.isFetching = true
            this.isUpdateInProgress = true

            this.notification = 'Updating configuration ...' + " @" + getDateAndTime()
            axios.post('http://localhost:8288/api/v1/conf/updateConfig', {
                "userId": this.userId,
                "email": this.emailId,
                "blogRepo": this.blogRepo,
                "siteRepo": this.siteRepo,
                "password": this.password
            })
                .then(response => {
                    this.notification = "Updated Successfully!" + " @" + getDateAndTime()
                    // reload the preview iframe
                    var previewFrame = document.getElementById('previewFrame');
                    previewFrame.src = previewFrame.src;
                    // Re-enable buttons
                    this.isFetching = false
                    this.isUpdateInProgress = false
                })
                .catch(error => {
                    this.notification = "Failed to update config!" + " @" + getDateAndTime()
                    
                    // Re-enable buttons
                    this.isFetching = false
                    this.isUpdateInProgress = false
                });
            
        }
    }

})
