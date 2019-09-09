<a href="https://cslab.aueb.gr"><img src="https://www.aueb.gr/press/logos/2_AUEB-white-HR.jpg" title="AUEB CSLab" alt="AUEB"></a>

> **Greek version [here](README.md)**

# Moniteur

Moniteur is a system for providing real time information during the academic year such as normal and examination schedule.

> Still WIP, but there is a working beta!

![](https://i.imgur.com/mu2FIDY.png)

---

## Contents

- [Technologies](#techonologies)
- [Building](#building)
- [Deployment](#deploment)
- [Documentation](#documentation)
- [Features](#featurs)
- [FAQ](#faq)
- [Team](#team)
- [License](#license)

---

## Techonologies

Moniteur takes advantage of state of the art technologies in order to complete this project.

Moniteur is consisted of three components, the backend , the frontend and a plugin for the backend.

### Backend

The backend has been written in [Go](https://golang.org/) and uses the following dependencies:

* [Labstack Echo](https://echo.labstack.com/)
* [JWT-Go](https://github.com/dgrijalva/jwt-go)
* [yaml](https://github.com/go-yaml/yaml)

For the function of the backend, we need to create a [Plugin](#Plugin).

---

### Plugin

The plugin has been written in [Go](https://golang.org/) as well. In this plugin we write all the logic so every organization that wants to use Moniteur writes a plugin.

In our case (AUEB) we use the following dependencies:

* [ldap](https://github.com/go-ldap/ldap)
* [xslsx-Go](https://github.com/tealeg/xlsx)

---

### Frontend

The frontend has been written in [VueJS](https://vuejs.org/).

---

## Building

In order to build Moniteur we need moniteur, the plugin and the frontend.

> **Building REQUIRES Linux operating system.**

> **Minimum version of Go 1.12.**

The steps for building are:

1. We install [golang](https://tecadmin.net/install-go-on-ubuntu/), [npm](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-16-04) using nvm method.
2. Clone το project ```git clone https://github.com/aueb-cslabs/moniteur.git```
3. Navigate to the project folder.
4. Edit [main.go](main.go) and add [here](https://github.com/aueb-cslabs/moniteur/blob/4bd80c4e78fdcf2af2a2569343c6261a5ed474bf/main.go#L48) the URL that the frontend is going to hit.
5. Edit [main.js](app/src/main.js) and change [here](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L56) the URL that the backend is going to use to access the API.
6. Edit the same file and change [here](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L57) the port of the backend that the API is.
7. Edit the same file and change [here](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L58) and change the logo.
8. Edit the same file and change [here](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L59) and add an extra logo if you need it.
9. Edit the same file and change [here](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L48) and change the locale to en.
10. Take advantage of the Makefile to build Moniteur. We analyze all the make function in the following table.

Make Command | Result
------------- | -------------------
.build-moniteur | Builds go project
.build-name-plugin | Builds go plugin
.build-go | Builds go project & plugin
.build-vue | Builds frontend
.build | Builds everything!

---

## Deployment

> **Deployment REQUIRES Linux operating system.**

After building we follow the steps bellow to deploy moniteur.

### Backend

Moniteur **requires** 2 configuration yml files in order to function.

* [config.yml](config.example.yml), contains the path of the plugin and the link of the exams program.
* [calendar.yml](calendar.example.yml), contains the academic calendar of the year.

In our case (AUEB) we need 2 more config files.

* [mapping.yml](mapping.example.yml), contains the mapping for every room for the needed english to greek translation of the URLs in order to take advantage of the existing API [AUEB Schedule Master](http://schedule.aueb.gr/).
* [ldap.yml](ldap.example.yml), contains the credentials of the ldap server and the users who have admin rights to Monituer.

After the creation of the files described above we follow below.

1. Navigate to Moniteur folder.
2. Create logs folder: ```mkdir logs```
3. Install tmux (if needed): ```sudo apt install tmux```
4. Create tmux container: ```tmux new -s moniteur```
5. Start Moniteur: ```./bin/moniteur```
6. Exit tmux container: ```Press Ctrl B + D```

### Frontend

> **Moniteur has been tested on Apache2 and Ubuntu 16.04! The following steps apply only on that system configuration!**

1. Install Apache2: ```sudo apt install apache2```
2. Stop the service: ```sudo systemctl stop apache2```
3. Copy the files from dist folder: ```sudo cp -R /somewhere/moniteur/app/dist/* /var/www/html/```
4. Enable overrides for Apache2: ```sudo a2enmod rewrite```
5. Navigate to html folder: ```cd /var/www/html/```
6. Create file named .htaccess: ```nano .htaccess```
7. Copy the following block: 
    ```
   #<IfModule mod_rewrite.c>
        RewriteEngine on
        RewriteCond %{REQUEST_FILENAME} -s [OR]
        RewriteCond %{REQUEST_FILENAME} -l [OR]
        RewriteCond %{REQUEST_FILENAME} -d
        RewriteRule ^.*$ - [NC,L]
   
        RewriteRule ^(.*) /index.html [NC, L]
   #</IfModule>
   ```
8. Save the file: ```Ctrl + X. After Y & enter```
9. Start Apache2: ```sudo systemctl start apache2```
10. Enable Apache2 on system start up: ```sudo systemctl enable apache2```

---

## Documentation

**Documentation is coming soon!**

---

## Features

Moniteur provides REST capabilities at users and the frontend displays all the information required per room.

The most simple way of using Moniteur for a normal user is to access the link: ```http://moniteur.url/roomID```

Moniteur automatically displays all the information for that specific room.

For Moniteur admins we provide creation of general and per room announcements and general comments. In order to access the admin panel you need to login at ```http://moniteur.url/admin```.

Automatically we provide POST, PUT and DELETE calls for the administration of Moniteur.

> **Detailed examples will be added then the projects comes at its first release!**

---

## FAQ

> **We will add FAQ in the following days, only if needed.**

---

## Team

Moniteur was created for the needs of AUEB's Computer Science Laboratories (CSLab).

The development of the project was conducted by the [Moniteur Team](https://github.com/orgs/aueb-cslabs/teams/moniteur).

---

## License

> **Coming Soon.**