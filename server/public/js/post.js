// vue component for post
var post = new Vue({
    el: '#new-post',
    data: {
        postname: '',
        status: ''
    },
    methods: {
        createPost: function () {
            // remove .md from postname
            var name = this.postname.replace('.md', '')
            this.status = ''
            if (name.length <1) {
                alert("Please enter post name")
                return
            }
            // call api to create new-post
            axios
                .get('api/v1/exec/newpost?name=' + name)
                .then(response => {
                    this.status = "successfully created post : " + name + " @" + getDateAndTime
                })
                .catch(error => {
                    this.status = "Failed to create post : " + name + ". Please check the name"
                    Console.log("Failed to created post :" + name)
                })

                this.postname = ''
        }
    }
})
