# vdock
#### Intro:
Vdock is a tool with an intergrated IDE for managing [hugo](https://gohugo.io/) site efficiently. 
<br>
It provides features like
* Integrated IDE (Cloud9)
* Live preview(changes)
* 1-click to create new post
* 1-click to publish final site to [github pages](https://pages.github.com/)
* snippets(markdown) to ease content writing

#### Run :
###### No Authentication & Deny Public Access
``` bash
$ docker run -td -p 8288:8288 -p 8286:8286 rchari/vdock
```
open : http://localhost:8288/

![Vdock Config page](pics/config.png)


<details><summary> More Options To Run</summary>
<p>

###### With Authentication & Deny Public Access
```bash
$ docker run -td -p 8288:8288 -p 8286:8286 -e AUTH=<username>:<password> rchari/vdock
```

###### No Authentication & Allow Public Access
```bash
$ docker run -td -p 8288:8288 -p 8286:8286 -e ALLOW_PUBLIC_ACCESS=YES rchari/vdock
```
###### With Authentication & Allow Public Access
```bash
$ docker run -td -e AUTH=<username>:<password> ALLOW_PUBLIC_ACCESS=YES \
    -p 8288:8288 -p 8286:8286 rchari/vdock
```
</p>
</details>

---
<br>

#### Create Github Repositories(Only once)
###### Step 1 : Fork below repository to your Github
> https://github.com/rchari66/hugo-sample-source

###### Step 2 : Create a new repository with name `<Github-userid>.github.io`
If your github-userid is `rchari66` then the repository name would be
> rchari66.github.io
---
<br>

#### Configure `vdock` in `config` page
<details><summary>Show config page</summary>
<p>

![Vdock Config page](pics/config.png)

</P>
</details>

###### Step 1 : Enter your github id and click `fetchRepositories` button
<details><summary>fetch Repositories</summary>
<p>

![Fetch Repositories](pics/fetch-repos.png)

</P>
</details>


###### Step 2 : Select repositories and Enter Other details.
> Select `hugo-sample-source` and `<Githuh-userid>.github.io` repos for `Blog Source:` and `Blog Destination:`

> Provide Github email and Github password and then Click on button `Update Configure`. 
This will create `blog-source` and `mySite` directries in the IDE.

<details><summary>Udate Config</summary>
<p>

![Udate configuration](pics/udate-config.png)


</P>
</details>

---
<br>

#### Configure config.toml(in `blog-source/config.toml`) for hugo
>  Make sure `publishDir` is set to `../mySite`.
```toml
publishDir = "../mySite"
```
> Update `baseURL` to `https://<Github-userid>.github.io`.
```toml
baseURL = "https://rchari66.github.io"
```
---
<br>

#### Publish your site.
> Simply goto `publish` page and click on `Publish Changes`(provide a commit message in textbox).

    Wait for one or two minutes and then open https://<Github-userid>.github.io to view your site.

In case the site in not available after couple of minutes then refer [here](https://help.github.com/articles/configuring-a-publishing-source-for-github-pages/)
for enabling github pages for your repository.
---
<br>

Refer : [How to Use vdock creating posts](https://techtaste.me/post/vdock-tool/)
