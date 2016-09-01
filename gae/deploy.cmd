goapp deploy default/app.yaml ^
&& goapp deploy redirect/app.yaml ^
&& goapp deploy api-service/app.yaml ^
&& appcfg.py -A uniform-shop update_dispatch . ^