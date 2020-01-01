[![Build Status](https://travis-ci.com/aueb-cslabs/moniteur.svg?branch=master)](https://travis-ci.com/aueb-cslabs/moniteur)

<a href="https://cslab.aueb.gr"><img src="https://www.aueb.gr/press/logos/2_AUEB-white-HR.jpg" title="AUEB CSLab" alt="AUEB"></a>

> **If you want to view the english version click [here](README_EN.md)**

# Moniteur

Το moniteur είναι ένα από σύστημα παρουσίασης του ακαδημαϊκού ημερολογίου καθώς και του εξαμηνιαίου και εξετασταίου προγράμματος.

> Still WIP, but there is a working beta!

![](https://i.imgur.com/mu2FIDY.png)

---

## Περιεχόμενα

- [Τεχνολογίες](#τεχνολογίες)
- [Κατασκευή](#κατασκευή)
- [Εγκατάσταση](#εγκατάσταση)
- [Documentation](#documentation)
- [Δυνατότητες](#δυνατότητες)
- [FAQ](#faq)
- [Team](#team)
- [License](#license)

---

## Τεχνολογίες

Το Moniteur έχει αξιοποιήσει τις τελευταίες τεχνολογίες που είναι διαθέσιμες για την δημιουργία αυτού του fullstack project.

Το Moniteur αποτελείται από τρία συστατικά συστατικά το backend, το frontend και το plugin.

### Backend

Το backend έχει υλοποιηθεί σε [Go](https://golang.org/) αξιοποιώντας αρκετά depedencies. Αναλυτικά τα εξής:

* [Labstack Echo](https://echo.labstack.com/)
* [JWT-Go](https://github.com/dgrijalva/jwt-go)
* [yaml](https://github.com/go-yaml/yaml)

Για την λειτουργία του απαιτείται η δημιουργία ενός [Plugin](#Plugin).

---

### Plugin

Το plugin έχει υλοποιηθεί και αυτό σε [Go](https://golang.org/). Ο κάθε οργανισμός που θέλει να αξιοποιήσει το Moniteur υλοποιεί όλη την λογική του σε αυτό το plugin.

Στην δική μας περίπτωση έχουμε χρησιμοποιήσει τα depedencies:

* [ldap](https://github.com/go-ldap/ldap)
* [xslsx-Go](https://github.com/tealeg/xlsx)

---

### Frontend

Το frontend έχει υλοποιηθεί σε [VueJS](https://vuejs.org/).

---

## Κατασκευή

Για την κατασκευή του project χρειαζόμαστε το moniteur, το plugin του moniteur καθώς και το frontend.


> **Για την κατασκευή του Moniteur ΑΠΑΙΤΕΙΤΑΙ Linux λειτουργικό σύστημα.**

> **Ελάχιστη έκδοση Go 1.12.**

Τα βήματα είναι τα εξής:

1. Εγκαθιστούμε [golang](https://tecadmin.net/install-go-on-ubuntu/), [npm](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-16-04) χρησιμοποιώντας την nvm μέθοδο.
2. Clone το project ```git clone https://github.com/aueb-cslabs/moniteur.git```
3. Πάμε στο φάκελο του project.
4. Επεξεργαζόμαστε το αρχείο [main.go](main.go) και προσθέτουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/4bd80c4e78fdcf2af2a2569343c6261a5ed474bf/main.go#L48) το URL που θα είναι το frontend μας.
5. Επεξεργαζόμαστε το αρχείο [main.js](app/src/main.js) και αλλάζουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L56) το URL που είναι το backend μας και θα χτυπάει το API.
6. Επεξεργαζόμαστε το ίδιο αρχείο και αλλάζουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L57) την πόρτα που είναι το API του backend μας.
7. Επεξεργαζόμαστε το ίδιο αρχείο και αλλάζουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L58) το logo μας.
8. Επεξεργαζόμαστε το ίδιο αρχείο και αλλάζουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L59) και προσθέτουμε προαιρετκά extra logo.
9. Επεξεργαζόμαστε το ίδιο αρχείο και αλλάζουμε [εδώ](https://github.com/aueb-cslabs/moniteur/blob/e4236dc8f72e7ebe71484e4f6a6f055acdc1e4bc/app/src/main.js#L48) και βάζουμε το locale στο el.
10. Αν θέλουμε να αλλάξουμε το background χρώμα, αλλάζουμε τον δεκαεξαδικό κωδικό [εδώ](https://github.com/aueb-cslabs/moniteur/blob/ea173109c674e28df6f66bcdc77142412b7475b9/app/src/main.js#L60).
11. Αν θέλουμε να αλλάξουμε το χρώμα του navbar, αλλάζουμε τον δεκαεξαδικό κωδικό [εδώ](https://github.com/aueb-cslabs/moniteur/blob/ea173109c674e28df6f66bcdc77142412b7475b9/app/src/main.js#L61).
12. Αξιοποιούμε το Makefile για να χτίσουμε το Moniteur. Παρακάτω αναλύονται οι εντολές του Makefile.

Make Command | Result
------------- | -------------------
.build-moniteur | Builds go project
.build-name-plugin | Builds go plugin
.build-go | Builds go project & plugin
.build-vue | Builds frontend
.build | Builds everything!

---

## Εγκατάσταση

> **Για την εγκατάσταση ΑΠΑΙΤΕΙΤΑΙ Linux λειτουργικό σύστημα.**

Μετά την κατασκευή έχουμε τα εξής βήματα για να εγκαταστήσουμε το Moniteur.

### Backend

Το Moniteur χρειάζεται 2 **υποχρεωτικά** αρχεία yml που περιέχουν το configuration του.

* [config.yml](config.example.yml), περιέχει την διαδρομή του plugin και του link του προγράμματος των εξετάσεων.
* [calendar.yml](calendar.example.yml), περιέχει το ακαδημαϊκό ημερολόγιο του έτους.

Στην περίπτωση του AUEB Plugin χρειαζόμαστε 2 ακόμα config αρχεία.

* [mapping.yml](mapping.example.yml), περιέχει το mapping των αιθουσών για την μετάφραση των αιθουσών για την λειτουργία με το [AUEB Schedule Master](http://schedule.aueb.gr/).
* [ldap.yml](ldap.example.yml), περιέχει τα στοιχεία του ldap server και τους χρήστες που έχουν πρόσβαση στο admin panel του Moniteur.

Μόλις δημιουργήσουμε τα παραπάνω αρχεία ακολουθούμε τα παρακάτω.

1. Πάμε στον φάκελο του Moniteur.
2. Δημιουργούμε τον φάκελο logs: ```mkdir logs```
3. Εγκαθιστούμε tmux: ```sudo apt install tmux```
4. Δημιουργούμε tmux container: ```tmux new -s moniteur```
5. Εκκινούμε το Moniteur: ```./bin/moniteur```
6. Έξοδος από tmux container: ```Press Ctrl B + D```

### Frontend

> **To Moniteur έχει δοκιμαστεί σε Apache2 και Ubuntu 16.04! Τα βήματα που ακολουθούν είναι για το συγκεκριμένο configuration!**

1. Εγκαθιστούμε Apache2: ```sudo apt install apache2```
2. Σταματάμε το service: ```sudo systemctl stop apache2```
3. Αντιγράφουμε τα αρχεία από τον dist φάκελο: ```sudo cp -R /somewhere/moniteur/app/dist/* /var/www/html/```
4. Ενεργοποιούμε overrides στον Apache2: ```sudo a2enmod rewrite```
5. Μετακινούμαστε στον φάκελο html: ```cd /var/www/html/```
6. Δημιουργούμε .htaccess αρχείο: ```nano .htaccess```
7. Αντιγράφουμε το επόμενο block: 
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
8. Αποθηκεύουμε το αρχείο: ```Ctrl + X. After Y & enter```
9. Εκκινούμε Apache2: ```sudo systemctl start apache2```
10. Ενεροποιούμε Apache2 στην εκκίνηση του λειτουργικού: ```sudo systemctl enable apache2```

---

## Documentation

**Έρχεται GoDoc που περιέχει όλο το documentation του κώδικα!**

---

## Δυνατότητες

Το Moniteur παρέχει REST δυνατότητες στους χρήστες του και το frontend προβάλει όλες τις πληροφορίες ανά αίθουσα.

Ο πιο απλός τρόπος χρήσης του Moniteur για έναν απλό χρήστη είναι να χτυπήσει το link: ```http://moniteur.url/roomID```

Το Moniteur αυτομάτως θα προβάλει όλες τις απαραίτητες πληροφορίες για αυτή την αίθουσα.

Για τους admins του Moniteur προσφέρονται οι δυνατότες δημιουργίας γενικών ανακοινώσεων, ανακοίνωση ανά αίθουσα και γενικοί σχολιασμοί. Για την πρόσβαση στο admin panel απαιτείται αυθεντικοποίηση στο ```http://moniteur.url/admin```.

Αμέσως προσφέρονται POST, PUT και DELETE κλήσεις για τον διαχειρισμό του Moniteur.

> **Αναλυτικότερα παραδείγματα θα προστεθούν μόλις ολοκληρωθεί το Moniteur!**

---

## FAQ

> **Θα προστεθεί FAQ τις ερχόμενες μέρες, αν χρειαστεί.**

---

## Team

Το Moniteur δημιουργήθηκε στα πλαίσια των αναγκών των εκπαιδευτικών εργαστιρίων πληροφορικής (CSLab) του AUEB.

Η ανάπτυξη πραγματοποιήθηκε από το [Moniteur Team](https://github.com/orgs/aueb-cslabs/teams/moniteur).

---

## License

> **Έρχεται σύντομα.**