
// vue component for config
var config = new Vue({
    el: '#set-config',
    data: {
        blogRepo: '',
        siteRepo: '',
        userId: '',
        email: 'raghavendrak555@gmail.com',
        password: '',
        repos_html:'',
        notification:''

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
            
            axios
                .get('api/v1/conf/getRepos?userId=' + name)
                .then(response => {
                    $("#set-config-repos").html(response.data)
                })
                .catch(error => {
                    alert("Either invalid git userid or empty repository list")
                    Console.log("Failed to get repository list")
                })

        },
        updateConfig: function () {
            // validate input length & content
            if (this.userId.length < 0 || this.password.length < 0) {
                alert("Please enter valid user details")
            }
            if (this.blogRepo.length < 0 || this.blogRepo.length < 0) {
                alert("Please enter valid repository details")
            }

            
            axios.post('http://localhost:8288/api/v1/conf/updateConfig', {
                "userId" : this.userId,
                "email"  : this.email,
                "blogRepo": this.blogRepo,
                "siteRepo": this.siteRepo,
                "password": this.password
            })
                .then(function (response) {
                    this.notification = "Updated Successfully!"
                })
                .catch(function (error) {
                    this.notification = "Failed to update config!"
                });
        }
    }

})
