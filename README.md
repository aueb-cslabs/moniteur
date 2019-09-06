Moniteur
=======

Spelled like a nice, regular-old *"monitor"*.

Still WIP, but there is a working beta!

There are two sides of to a story always, so let's start with the back-end.

###Backend

Moniteur is written in golang. Minimum version required is 1.11.

Moniteur has 3 config files, two of them are mandatory.

* config.yml, This file contains the plugin location and the exams link.
* calendar.yml, This file contains the academic calendar
* mapping.yml, This file is *NOT* required. It contains the room mapping for english-greek link translation.

There are example files in the repo.

###Frontend

Moniteur is written in VueJS.

For the frontend to work properly you **need to have overrides enabled**.

###Building

Install latest golang and npm or yarn.

Steps for building moniteur:

1. Navigate to project folder.
2. Rename all the api links in the Vue project (app folder) from localhost to the URL that moniteur is gonna hit.
3. Add that URL to CORS config in main.go.
4. Build the project using make.
 
Make use of the makefile to build the project.

Make Command | Result
------------- | -------------------
.build-moniteur | Builds go project
.build-name-plugin | Builds go plugin
.build-go | Builds go project & plugin
.build-vue | Builds frontend
.build | Builds everything!

###Deployment

**DEPLOYMENT ONLY ON LINUX SYSTEMS**

After building everything, these are the steps in order to deploy Moniteur successfully.

For the backend:

1. Navigate to project folder
2. Create logs folder: ```mkdir logs```
3. Install tmux: ```sudo apt install tmux```
4. Create container: ```tmux new -s moniteur```
5. Start moniteur backend: ```./bin/moniteur```
6. Exit container: ```Press Ctrl B + D```

For the frontend:

Moniteur frontend has been tested on Apache2 and Ubuntu 16.04 for the time being.

So the steps that follow apply only to Apache2 and Ubuntu 16.04.

1. Install Apache2: ```sudo apt install apache2```
2. Stop the service if running: ```sudo systemctl stop apache2```
3. Copy all the files from dist folder: ```sudo cp -R /somewhere/moniteur/app/dist/* /var/www/html/```
4. Enable overrides on Apache2: ```sudo a2enmod rewrite```
5. Navigate over the html folder: ```cd /var/www/html/```
6. Create .htaccess file: ```nano .htaccess```
7. Paste the following: 
    ```
   #<IfModule mod_rewrite.c>
        RewriteEngine on
        RewriteCond %{REQUEST_FILENAME} -s [OR]
        RewriteCond %{REQUEST_FILENAME} -l [OR]
        RewriteCond %{REQUEST_FILENAME} -d
        RewriteRule ^.*$ - [NC,L]
   
        RewriteRule ^(.*) /index.html [LC, L]
   #</IfModule>
   ```
8. Start Apache2: ```sudo systemctl start apache2```
9. Enable Apache2 on startup: ```sudo systemctl enable apache2```

###How to use

Open up any browser and enter: ```http://moniteur.url/room```

###AUEB Case of Use

Currently AUEB CSLabs have installed Raspberry Pi 3 and monitors.

Every lab displays its program, announcements and exams from Moniteur using the Pi.