{% for up in upstreams %}
upstream {{up.Name}} {
    least_conn;
    {% for target in up.Targets %}server {{target.Addr}} weight={{target.Weight}} max_fails=2 fail_timeout=6s;
    {% endfor %}
    keepalive 128;
}
{% endfor %}

server {
    server_name  {{server.Domain}};
    listen 80;
    listen 443 ssl;

    ssl_certificate cert/autohome.crt;
    ssl_certificate_key cert/autohome.key;

    ssl_ciphers "EECDH+AESGCM:EDH+AESGCM:ECDHE-RSA-AES128-GCM-SHA256:AES256+EECDH:DHE-RSA-AES128-GCM-SHA256:AES256+EDH:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES128-SHA256:DHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES256-GCM-SHA384:AES128-GCM-SHA256:AES256-SHA256:AES128-SHA256:AES256-SHA:AES128-SHA:DES-CBC3-SHA:HIGH:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4";
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;

    index index.aspx index.shtml index.html;

    ssl_prefer_server_ciphers on;
    keepalive_timeout 65;
    ssl_session_cache shared:SSL:512m;
    ssl_buffer_size 1400;
    ssl_session_timeout 30m;


    {% for location in server.Locations %}
    location /{{location.ProjectName}}/(.*)$ {
        proxy_set_header Host {{location.ProxyHost}};
        proxy_pass http://{{location.Upstream}}/$1$is_args$args;
    }
    {% endfor %}
}



