
// vue component for post
var post = new Vue({
    el: '#new-post',
    data: {
        postname: '',
        status: '',
        isCreateInProgress: false
    },
    methods: {
        createPost: function () {
            // remove .md and replace space with '-'
            var name = this.postname.replace('.md', '').replace(' ', '-')
            this.status = ''
            if (name.length <1) {
                alert("Please enter post name")
                return
            }
            this.isCreateInProgress = true
            // call api to create new-post
            axios
                .get('api/v1/exec/newpost?name=' + name)
                .then(response => {
                    this.status = "successfully created post : " + name + " @" + getDateAndTime()
                    this.isCreateInProgress = false
                })
                .catch(error => {
                    this.status = "Failed to create post : " + name + ". Please check the name"
                    this.isCreateInProgress = false
                })

                this.postname = ''
        }
    }
})
