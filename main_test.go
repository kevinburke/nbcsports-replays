package main

import (
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	r := strings.NewReader(page)
	resp, err := getLinks(r)
	if err != nil {
		t.Fatal(err)
	}
	l := len(resp.LiveExtra.Events)
	if l != 16 {
		t.Fatalf("expected length 16, got %d", l)
	}
	network := resp.LiveExtra.Events[0].Network
	if network != "NBCSN" {
		t.Fatalf("expected network NBCSN, got %s", network)
	}
	title := resp.LiveExtra.Events[0].Title
	if title != "Swansea City v. Liverpool" {
		t.Fatalf("expected title Liverpool vs Swansea, got %v", title)
	}
}

func TestParseHTML2(t *testing.T) {
	r := strings.NewReader(page)
	events, err := getEvents(r)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) == 0 {
		t.Fatal("expected to get some events, got 0")
	}
	if events[0].Title == "" {
		t.Fatal("expected events[0].Title to be non-empty, got an empty string")
	}
}

const page = `
<!DOCTYPE html>
<html>
<head>
  <!--[if IE]><![endif]-->
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-180x180.png" rel="apple-touch-icon-precomposed" sizes="180x180" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-144x144.png" rel="apple-touch-icon-precomposed" sizes="144x144" />
<link rel="shortcut icon" href="http://www.nbcsports.com/sites/nbcsports.com/files/favicon.ico" type="image/vnd.microsoft.icon" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-57x57.png" rel="apple-touch-icon-precomposed" sizes="57x57" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-72x72.png" rel="apple-touch-icon-precomposed" sizes="72x72" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-152x152.png" rel="apple-touch-icon-precomposed" sizes="152x152" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-114x114.png" rel="apple-touch-icon-precomposed" sizes="114x114" />
<link href="/sites/all/themes/custom/nbcsports/img/apple-touch-icons/nbcsports/ios/apple-touch-icon-120x120.png" rel="apple-touch-icon-precomposed" sizes="120x120" />
<meta name="tp:EnableExternalController" content="true" />
<meta name="description" content="NBCSports.com provides up-to-the-minute coverage of sports, including live streaming, on-demand video and text-based content." />
<meta name="generator" content="Drupal 7 (http://drupal.org)" />
<link rel="canonical" href="http://www.nbcsports.com/live" />
<link rel="shortlink" href="http://www.nbcsports.com/live" />
<meta property="og:site_name" content="NBC Sports" />
<meta property="og:type" content="article" />
<meta property="og:url" content="http://www.nbcsports.com/live" />
<meta property="og:title" content="Live and Upcoming" />
<meta name="twitter:card" content="summary" />
<meta name="twitter:url" content="http://www.nbcsports.com/live" />
<meta name="twitter:title" content="Live and Upcoming" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge"><script type="text/javascript">(window.NREUM||(NREUM={})).loader_config={xpid:"VwIFV1ZRGwIFU1NSBAUC"};window.NREUM||(NREUM={}),__nr_require=function(t,e,n){function r(n){if(!e[n]){var o=e[n]={exports:{}};t[n][0].call(o.exports,function(e){var o=t[n][1][e];return r(o||e)},o,o.exports)}return e[n].exports}if("function"==typeof __nr_require)return __nr_require;for(var o=0;o<n.length;o++)r(n[o]);return r}({1:[function(t,e,n){function r(t){try{s.console&&console.log(t)}catch(e){}}var o,i=t("ee"),a=t(15),s={};try{o=localStorage.getItem("__nr_flags").split(","),console&&"function"==typeof console.log&&(s.console=!0,o.indexOf("dev")!==-1&&(s.dev=!0),o.indexOf("nr_dev")!==-1&&(s.nrDev=!0))}catch(c){}s.nrDev&&i.on("internal-error",function(t){r(t.stack)}),s.dev&&i.on("fn-err",function(t,e,n){r(n.stack)}),s.dev&&(r("NR AGENT IN DEVELOPMENT MODE"),r("flags: "+a(s,function(t,e){return t}).join(", ")))},{}],2:[function(t,e,n){function r(t,e,n,r,o){try{d?d-=1:i("err",[o||new UncaughtException(t,e,n)])}catch(s){try{i("ierr",[s,(new Date).getTime(),!0])}catch(c){}}return"function"==typeof f&&f.apply(this,a(arguments))}function UncaughtException(t,e,n){this.message=t||"Uncaught error with no additional information",this.sourceURL=e,this.line=n}function o(t){i("err",[t,(new Date).getTime()])}var i=t("handle"),a=t(16),s=t("ee"),c=t("loader"),f=window.onerror,u=!1,d=0;c.features.err=!0,t(1),window.onerror=r;try{throw new Error}catch(l){"stack"in l&&(t(8),t(7),"addEventListener"in window&&t(5),c.xhrWrappable&&t(9),u=!0)}s.on("fn-start",function(t,e,n){u&&(d+=1)}),s.on("fn-err",function(t,e,n){u&&(this.thrown=!0,o(n))}),s.on("fn-end",function(){u&&!this.thrown&&d>0&&(d-=1)}),s.on("internal-error",function(t){i("ierr",[t,(new Date).getTime(),!0])})},{}],3:[function(t,e,n){t("loader").features.ins=!0},{}],4:[function(t,e,n){function r(t){}if(window.performance&&window.performance.timing&&window.performance.getEntriesByType){var o=t("ee"),i=t("handle"),a=t(8),s=t(7),c="learResourceTimings",f="addEventListener",u="resourcetimingbufferfull",d="bstResource",l="resource",p="-start",h="-end",m="fn"+p,w="fn"+h,v="bstTimer",y="pushState";t("loader").features.stn=!0,t(6);var g=NREUM.o.EV;o.on(m,function(t,e){var n=t[0];n instanceof g&&(this.bstStart=Date.now())}),o.on(w,function(t,e){var n=t[0];n instanceof g&&i("bst",[n,e,this.bstStart,Date.now()])}),a.on(m,function(t,e,n){this.bstStart=Date.now(),this.bstType=n}),a.on(w,function(t,e){i(v,[e,this.bstStart,Date.now(),this.bstType])}),s.on(m,function(){this.bstStart=Date.now()}),s.on(w,function(t,e){i(v,[e,this.bstStart,Date.now(),"requestAnimationFrame"])}),o.on(y+p,function(t){this.time=Date.now(),this.startPath=location.pathname+location.hash}),o.on(y+h,function(t){i("bstHist",[location.pathname+location.hash,this.startPath,this.time])}),f in window.performance&&(window.performance["c"+c]?window.performance[f](u,function(t){i(d,[window.performance.getEntriesByType(l)]),window.performance["c"+c]()},!1):window.performance[f]("webkit"+u,function(t){i(d,[window.performance.getEntriesByType(l)]),window.performance["webkitC"+c]()},!1)),document[f]("scroll",r,!1),document[f]("keypress",r,!1),document[f]("click",r,!1)}},{}],5:[function(t,e,n){function r(t){for(var e=t;e&&!e.hasOwnProperty(u);)e=Object.getPrototypeOf(e);e&&o(e)}function o(t){s.inPlace(t,[u,d],"-",i)}function i(t,e){return t[1]}var a=t("ee").get("events"),s=t(17)(a),c=t("gos"),f=XMLHttpRequest,u="addEventListener",d="removeEventListener";e.exports=a,"getPrototypeOf"in Object?(r(document),r(window),r(f.prototype)):f.prototype.hasOwnProperty(u)&&(o(window),o(f.prototype)),a.on(u+"-start",function(t,e){if(t[1]){var n=t[1];if("function"==typeof n){var r=c(n,"nr@wrapped",function(){return s(n,"fn-",null,n.name||"anonymous")});this.wrapped=t[1]=r}else"function"==typeof n.handleEvent&&s.inPlace(n,["handleEvent"],"fn-")}}),a.on(d+"-start",function(t){var e=this.wrapped;e&&(t[1]=e)})},{}],6:[function(t,e,n){var r=t("ee").get("history"),o=t(17)(r);e.exports=r,o.inPlace(window.history,["pushState","replaceState"],"-")},{}],7:[function(t,e,n){var r=t("ee").get("raf"),o=t(17)(r),i="equestAnimationFrame";e.exports=r,o.inPlace(window,["r"+i,"mozR"+i,"webkitR"+i,"msR"+i],"raf-"),r.on("raf-start",function(t){t[0]=o(t[0],"fn-")})},{}],8:[function(t,e,n){function r(t,e,n){t[0]=a(t[0],"fn-",null,n)}function o(t,e,n){this.method=n,this.timerDuration="number"==typeof t[1]?t[1]:0,t[0]=a(t[0],"fn-",this,n)}var i=t("ee").get("timer"),a=t(17)(i),s="setTimeout",c="setInterval",f="clearTimeout",u="-start",d="-";e.exports=i,a.inPlace(window,[s,"setImmediate"],s+d),a.inPlace(window,[c],c+d),a.inPlace(window,[f,"clearImmediate"],f+d),i.on(c+u,r),i.on(s+u,o)},{}],9:[function(t,e,n){function r(t,e){d.inPlace(e,["onreadystatechange"],"fn-",s)}function o(){var t=this,e=u.context(t);t.readyState>3&&!e.resolved&&(e.resolved=!0,u.emit("xhr-resolved",[],t)),d.inPlace(t,w,"fn-",s)}function i(t){v.push(t),h&&(g=-g,b.data=g)}function a(){for(var t=0;t<v.length;t++)r([],v[t]);v.length&&(v=[])}function s(t,e){return e}function c(t,e){for(var n in t)e[n]=t[n];return e}t(5);var f=t("ee"),u=f.get("xhr"),d=t(17)(u),l=NREUM.o,p=l.XHR,h=l.MO,m="readystatechange",w=["onload","onerror","onabort","onloadstart","onloadend","onprogress","ontimeout"],v=[];e.exports=u;var y=window.XMLHttpRequest=function(t){var e=new p(t);try{u.emit("new-xhr",[e],e),e.addEventListener(m,o,!1)}catch(n){try{u.emit("internal-error",[n])}catch(r){}}return e};if(c(p,y),y.prototype=p.prototype,d.inPlace(y.prototype,["open","send"],"-xhr-",s),u.on("send-xhr-start",function(t,e){r(t,e),i(e)}),u.on("open-xhr-start",r),h){var g=1,b=document.createTextNode(g);new h(a).observe(b,{characterData:!0})}else f.on("fn-end",function(t){t[0]&&t[0].type===m||a()})},{}],10:[function(t,e,n){function r(t){var e=this.params,n=this.metrics;if(!this.ended){this.ended=!0;for(var r=0;r<d;r++)t.removeEventListener(u[r],this.listener,!1);if(!e.aborted){if(n.duration=(new Date).getTime()-this.startTime,4===t.readyState){e.status=t.status;var i=o(t,this.lastSize);if(i&&(n.rxSize=i),this.sameOrigin){var a=t.getResponseHeader("X-NewRelic-App-Data");a&&(e.cat=a.split(", ").pop())}}else e.status=0;n.cbTime=this.cbTime,f.emit("xhr-done",[t],t),s("xhr",[e,n,this.startTime])}}}function o(t,e){var n=t.responseType;if("json"===n&&null!==e)return e;var r="arraybuffer"===n||"blob"===n||"json"===n?t.response:t.responseText;return h(r)}function i(t,e){var n=c(e),r=t.params;r.host=n.hostname+":"+n.port,r.pathname=n.pathname,t.sameOrigin=n.sameOrigin}var a=t("loader");if(a.xhrWrappable){var s=t("handle"),c=t(11),f=t("ee"),u=["load","error","abort","timeout"],d=u.length,l=t("id"),p=t(14),h=t(13),m=window.XMLHttpRequest;a.features.xhr=!0,t(9),f.on("new-xhr",function(t){var e=this;e.totalCbs=0,e.called=0,e.cbTime=0,e.end=r,e.ended=!1,e.xhrGuids={},e.lastSize=null,p&&(p>34||p<10)||window.opera||t.addEventListener("progress",function(t){e.lastSize=t.loaded},!1)}),f.on("open-xhr-start",function(t){this.params={method:t[0]},i(this,t[1]),this.metrics={}}),f.on("open-xhr-end",function(t,e){"loader_config"in NREUM&&"xpid"in NREUM.loader_config&&this.sameOrigin&&e.setRequestHeader("X-NewRelic-ID",NREUM.loader_config.xpid)}),f.on("send-xhr-start",function(t,e){var n=this.metrics,r=t[0],o=this;if(n&&r){var i=h(r);i&&(n.txSize=i)}this.startTime=(new Date).getTime(),this.listener=function(t){try{"abort"===t.type&&(o.params.aborted=!0),("load"!==t.type||o.called===o.totalCbs&&(o.onloadCalled||"function"!=typeof e.onload))&&o.end(e)}catch(n){try{f.emit("internal-error",[n])}catch(r){}}};for(var a=0;a<d;a++)e.addEventListener(u[a],this.listener,!1)}),f.on("xhr-cb-time",function(t,e,n){this.cbTime+=t,e?this.onloadCalled=!0:this.called+=1,this.called!==this.totalCbs||!this.onloadCalled&&"function"==typeof n.onload||this.end(n)}),f.on("xhr-load-added",function(t,e){var n=""+l(t)+!!e;this.xhrGuids&&!this.xhrGuids[n]&&(this.xhrGuids[n]=!0,this.totalCbs+=1)}),f.on("xhr-load-removed",function(t,e){var n=""+l(t)+!!e;this.xhrGuids&&this.xhrGuids[n]&&(delete this.xhrGuids[n],this.totalCbs-=1)}),f.on("addEventListener-end",function(t,e){e instanceof m&&"load"===t[0]&&f.emit("xhr-load-added",[t[1],t[2]],e)}),f.on("removeEventListener-end",function(t,e){e instanceof m&&"load"===t[0]&&f.emit("xhr-load-removed",[t[1],t[2]],e)}),f.on("fn-start",function(t,e,n){e instanceof m&&("onload"===n&&(this.onload=!0),("load"===(t[0]&&t[0].type)||this.onload)&&(this.xhrCbStart=(new Date).getTime()))}),f.on("fn-end",function(t,e){this.xhrCbStart&&f.emit("xhr-cb-time",[(new Date).getTime()-this.xhrCbStart,this.onload,e],e)})}},{}],11:[function(t,e,n){e.exports=function(t){var e=document.createElement("a"),n=window.location,r={};e.href=t,r.port=e.port;var o=e.href.split("://");!r.port&&o[1]&&(r.port=o[1].split("/")[0].split("@").pop().split(":")[1]),r.port&&"0"!==r.port||(r.port="https"===o[0]?"443":"80"),r.hostname=e.hostname||n.hostname,r.pathname=e.pathname,r.protocol=o[0],"/"!==r.pathname.charAt(0)&&(r.pathname="/"+r.pathname);var i=!e.protocol||":"===e.protocol||e.protocol===n.protocol,a=e.hostname===document.domain&&e.port===n.port;return r.sameOrigin=i&&(!e.hostname||a),r}},{}],12:[function(t,e,n){function r(){}function o(t,e,n){return function(){return i(t,[(new Date).getTime()].concat(s(arguments)),e?null:this,n),e?void 0:this}}var i=t("handle"),a=t(15),s=t(16),c=t("ee").get("tracer"),f=NREUM;"undefined"==typeof window.newrelic&&(newrelic=f);var u=["setPageViewName","setCustomAttribute","setErrorHandler","finished","addToTrace","inlineHit"],d="api-",l=d+"ixn-";a(u,function(t,e){f[e]=o(d+e,!0,"api")}),f.addPageAction=o(d+"addPageAction",!0),e.exports=newrelic,f.interaction=function(){return(new r).get()};var p=r.prototype={createTracer:function(t,e){var n={},r=this,o="function"==typeof e;return i(l+"tracer",[Date.now(),t,n],r),function(){if(c.emit((o?"":"no-")+"fn-start",[Date.now(),r,o],n),o)try{return e.apply(this,arguments)}finally{c.emit("fn-end",[Date.now()],n)}}}};a("setName,setAttribute,save,ignore,onEnd,getContext,end,get".split(","),function(t,e){p[e]=o(l+e)}),newrelic.noticeError=function(t){"string"==typeof t&&(t=new Error(t)),i("err",[t,(new Date).getTime()])}},{}],13:[function(t,e,n){e.exports=function(t){if("string"==typeof t&&t.length)return t.length;if("object"==typeof t){if("undefined"!=typeof ArrayBuffer&&t instanceof ArrayBuffer&&t.byteLength)return t.byteLength;if("undefined"!=typeof Blob&&t instanceof Blob&&t.size)return t.size;if(!("undefined"!=typeof FormData&&t instanceof FormData))try{return JSON.stringify(t).length}catch(e){return}}}},{}],14:[function(t,e,n){var r=0,o=navigator.userAgent.match(/Firefox[\/\s](\d+\.\d+)/);o&&(r=+o[1]),e.exports=r},{}],15:[function(t,e,n){function r(t,e){var n=[],r="",i=0;for(r in t)o.call(t,r)&&(n[i]=e(r,t[r]),i+=1);return n}var o=Object.prototype.hasOwnProperty;e.exports=r},{}],16:[function(t,e,n){function r(t,e,n){e||(e=0),"undefined"==typeof n&&(n=t?t.length:0);for(var r=-1,o=n-e||0,i=Array(o<0?0:o);++r<o;)i[r]=t[e+r];return i}e.exports=r},{}],17:[function(t,e,n){function r(t){return!(t&&"function"==typeof t&&t.apply&&!t[a])}var o=t("ee"),i=t(16),a="nr@original",s=Object.prototype.hasOwnProperty,c=!1;e.exports=function(t){function e(t,e,n,o){function nrWrapper(){var r,a,s,c;try{a=this,r=i(arguments),s="function"==typeof n?n(r,a):n||{}}catch(u){d([u,"",[r,a,o],s])}f(e+"start",[r,a,o],s);try{return c=t.apply(a,r)}catch(l){throw f(e+"err",[r,a,l],s),l}finally{f(e+"end",[r,a,c],s)}}return r(t)?t:(e||(e=""),nrWrapper[a]=t,u(t,nrWrapper),nrWrapper)}function n(t,n,o,i){o||(o="");var a,s,c,f="-"===o.charAt(0);for(c=0;c<n.length;c++)s=n[c],a=t[s],r(a)||(t[s]=e(a,f?s+o:o,i,s))}function f(e,n,r){if(!c){c=!0;try{t.emit(e,n,r)}catch(o){d([o,e,n,r])}c=!1}}function u(t,e){if(Object.defineProperty&&Object.keys)try{var n=Object.keys(t);return n.forEach(function(n){Object.defineProperty(e,n,{get:function(){return t[n]},set:function(e){return t[n]=e,e}})}),e}catch(r){d([r])}for(var o in t)s.call(t,o)&&(e[o]=t[o]);return e}function d(e){try{t.emit("internal-error",e)}catch(n){}}return t||(t=o),e.inPlace=n,e.flag=a,e}},{}],ee:[function(t,e,n){function r(){}function o(t){function e(t){return t&&t instanceof r?t:t?s(t,a,i):i()}function n(n,r,o){t&&t(n,r,o);for(var i=e(o),a=l(n),s=a.length,c=0;c<s;c++)a[c].apply(i,r);var u=f[w[n]];return u&&u.push([v,n,r,i]),i}function d(t,e){m[t]=l(t).concat(e)}function l(t){return m[t]||[]}function p(t){return u[t]=u[t]||o(n)}function h(t,e){c(t,function(t,n){e=e||"feature",w[n]=e,e in f||(f[e]=[])})}var m={},w={},v={on:d,emit:n,get:p,listeners:l,context:e,buffer:h};return v}function i(){return new r}var a="nr@context",s=t("gos"),c=t(15),f={},u={},d=e.exports=o();d.backlog=f},{}],gos:[function(t,e,n){function r(t,e,n){if(o.call(t,e))return t[e];var r=n();if(Object.defineProperty&&Object.keys)try{return Object.defineProperty(t,e,{value:r,writable:!0,enumerable:!1}),r}catch(i){}return t[e]=r,r}var o=Object.prototype.hasOwnProperty;e.exports=r},{}],handle:[function(t,e,n){function r(t,e,n,r){o.buffer([t],r),o.emit(t,e,n)}var o=t("ee").get("handle");e.exports=r,r.ee=o},{}],id:[function(t,e,n){function r(t){var e=typeof t;return!t||"object"!==e&&"function"!==e?-1:t===window?0:a(t,i,function(){return o++})}var o=1,i="nr@id",a=t("gos");e.exports=r},{}],loader:[function(t,e,n){function r(){if(!g++){var t=y.info=NREUM.info,e=u.getElementsByTagName("script")[0];if(t&&t.licenseKey&&t.applicationID&&e){c(w,function(e,n){t[e]||(t[e]=n)});var n="https"===m.split(":")[0]||t.sslForHttp;y.proto=n?"https://":"http://",s("mark",["onload",a()],null,"api");var r=u.createElement("script");r.src=y.proto+t.agent,e.parentNode.insertBefore(r,e)}}}function o(){"complete"===u.readyState&&i()}function i(){s("mark",["domContent",a()],null,"api")}function a(){return(new Date).getTime()}var s=t("handle"),c=t(15),f=window,u=f.document,d="addEventListener",l="attachEvent",p=f.XMLHttpRequest,h=p&&p.prototype;NREUM.o={ST:setTimeout,CT:clearTimeout,XHR:p,REQ:f.Request,EV:f.Event,PR:f.Promise,MO:f.MutationObserver},t(12);var m=""+location,w={beacon:"bam.nr-data.net",errorBeacon:"bam.nr-data.net",agent:"js-agent.newrelic.com/nr-974.min.js"},v=p&&h&&h[d]&&!/CriOS/.test(navigator.userAgent),y=e.exports={offset:a(),origin:m,features:{},xhrWrappable:v};u[d]?(u[d]("DOMContentLoaded",i,!1),f[d]("load",r,!1)):(u[l]("onreadystatechange",o),f[l]("onload",r)),s("mark",["firstbyte",a()],null,"api");var g=0},{}]},{},["loader",2,10,4,3]);</script>
  <meta name=viewport content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
  <title>Live and Upcoming | NBC Sports</title>
  <!-- Don't remove this. We are loading base file at first -->
  <style id="base-css">html{font-family:sans-serif;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}body{margin:0}article,aside,details,figcaption,figure,footer,header,main,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block;vertical-align:baseline}audio:not([controls]){display:none;height:0}[hidden],template{display:none}a{background:transparent}a:active,a:hover{outline:0}abbr[title]{border-bottom:1px dotted}b,strong{font-weight:bold}dfn{font-style:italic}h1{font-size:2em;margin:0.67em 0}mark{background:#ff0;color:#000}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:baseline}sup{top:-0.5em}sub{bottom:-0.25em}img{border:0}svg:not(:root){overflow:hidden}figure{margin:1em 40px}hr{box-sizing:content-box;height:0}pre{overflow:auto}code,kbd,pre,samp{font-family:monospace, monospace;font-size:1em}button,input,optgroup,select,textarea{color:inherit;font:inherit;margin:0}button{overflow:visible}button,select{text-transform:none}button,html input[type="button"],input[type="reset"],input[type="submit"]{-webkit-appearance:button;cursor:pointer}button[disabled],html input[disabled]{cursor:default}button::-moz-focus-inner,input::-moz-focus-inner{border:0;padding:0}input{line-height:normal}input[type="checkbox"],input[type="radio"]{box-sizing:border-box;padding:0}input[type="number"]::-webkit-inner-spin-button,input[type="number"]::-webkit-outer-spin-button{height:auto}input[type="search"]{-webkit-appearance:textfield;box-sizing:content-box}input[type="search"]::-webkit-search-cancel-button,input[type="search"]::-webkit-search-decoration{-webkit-appearance:none}fieldset{border:1px solid #c0c0c0;margin:0 2px;padding:0.35em 0.625em 0.75em}legend{border:0;padding:0}textarea{overflow:auto}optgroup{font-weight:bold}table{border-collapse:collapse;border-spacing:0}td,th{padding:0}html{font-family:Arial, Helvetica, sans-serif}a{text-decoration:none}h1,.alpha,h2,.beta,h3,.gamma,h4,.delta,h5,.epsilon,h6,.zeta{font-family:Arial, Helvetica, sans-serif;text-transform:uppercase}h1 a,h1 a:visited,.alpha a,.alpha a:visited,h2 a,h2 a:visited,.beta a,.beta a:visited,h3 a,h3 a:visited,.gamma a,.gamma a:visited,h4 a,h4 a:visited,.delta a,.delta a:visited,h5 a,h5 a:visited,.epsilon a,.epsilon a:visited,h6 a,h6 a:visited,.zeta a,.zeta a:visited{text-decoration:none}h1,.alpha{font-family:"BebasNeue", Arial, Helvetica, sans-serif;font-size:40px;font-size:2.5rem;font-weight:400}h2,.beta{font-size:34px;font-size:2.125rem}h3,.gamma{font-size:21px;font-size:1.312rem}h4,.delta{font-size:18px;font-size:1.125rem}h5,.epsilon{font-size:16px;font-size:1rem}h6,.zeta{font-size:16px;font-size:1rem}p{margin:0 0 1em}blockquote{padding:0 1.5em;font-style:italic}blockquote p{margin:0}q{font-style:italic}mark,.marker{padding:0.1em 0.5em;background:#666;color:#fff;font-size:12px;font-size:0.75rem;text-transform:capitalize}::-webkit-input-placeholder{color:#fff}:-moz-placeholder{color:#fff;opacity:1}::-moz-placeholder{color:#fff;opacity:1}:-ms-input-placeholder{color:#fff}img{max-width:100%;height:auto;vertical-align:top}.media-box iframe{width:100%;height:auto}@media (min-width: 1260px){.desktop-large_hidden {display: none !important;}.desktop-large_visible{display: block !important;}}@media (min-width: 1024px) and (max-width: 1260px){.desktop-small_hidden{display: none !important;}.desktop-small_visible{display: block !important;}}@media (min-width: 320px) and (max-width: 767px){.mobile_hidden{display:none !important}.mobile_visible{display:block !important}}@media (min-width: 768px) and (max-width: 1023px){.tablet_hidden{display:none !important}.tablet_visible{display:block !important}}</style>
  <link type="text/css" rel="stylesheet" href="//www.nbcsports.com/sites/nbcsports.com/files/advagg_css/css__LqXDLOEdu4FjANrDpUp-e9hLLmTMujC_lNNowZP2pPk__4j7vksssq8uPKxjky87FDjEQeTo92tKxLYjUQ1GU1K8__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.css" media="all" />
<link type="text/css" rel="stylesheet" href="//www.nbcsports.com/sites/nbcsports.com/files/advagg_css/css__dprMNsqUqFAs7_VAlnYRrJJSvSuiUgvwqeZPECqXJf0__gUsIyWPv3uxAPYj4lFwdjegOWGB39uf7sdob124eJzE__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.css" media="all" />
<link type="text/css" rel="stylesheet" href="//www.nbcsports.com/sites/nbcsports.com/files/advagg_css/css__zM-ZkzO8EIWM9uffpGgpMcFMqbXH0ScoAFumCfVGEig__Z4wgbr7JK4d92mgb7Z64Hi1clWQkP7D-3hhk1l89srw__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.css" media="all" />
<link type="text/css" rel="stylesheet" href="//www.nbcsports.com/sites/nbcsports.com/files/advagg_css/css__1bJpa14GvFoLtvJc7hP1U0q_aI1qELsyfCYUZ38fxcg__85TtC1jiWE2gS68RUoOYA-sCq3eO0IXPgF1mcOpXE-g__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.css" media="all" />
  <!-- Script for catch errors on Js error testing.
       It should goes before Drupal scripts array and any placed script -->
  <script>
    window.jsErrors = [];
    window.onerror = function(errorMessage) {
      window.jsErrors[window.jsErrors.length] = errorMessage;
    };
  </script>
  <script type="text/javascript" src="//www.nbcsports.com/sites/nbcsports.com/files/advagg_js/js__iX0QvshDzCXgZggnpaRfw_C3u6X_UAjCBBEeQnQgcSo__fUgnjqmjX9LyzvvvkU3MMW81yEbwf96FW79DFYknIho__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.js"></script>
<script type="text/javascript">
<!--//--><![CDATA[//><!--
document.createElement( "picture" );
//--><!]]>
</script>
<script type="text/javascript" src="//www.nbcsports.com/sites/nbcsports.com/files/advagg_js/js__CNkwMrd76LbgBpe8m-LAi3EqnCzaRNMBlyiVQY-5YLU__y3uJNjMrk-UUz9riUnY4jl6wIVGTcQKJb32xDfioJfc__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.js" async="async"></script>
<script type="text/javascript" src="//www.nbcsports.com/sites/nbcsports.com/files/advagg_js/js__gEqAYrL4rFsP2KK2VBkiScFG09LfanuGZH-R_oQI6Hs__3X6GH467jjYfIOSEp4XeyFTjEnTkjVtwCMXI_x_VuZY__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.js"></script>
<script type="text/javascript">
<!--//--><![CDATA[//><!--
var mpscall = {"site":"nbcsports-web","path":"\/live","title":"Live and Upcoming","is_content":"0","type":"other"}
var mpsopts = {"host":"mps.nbcuni.com","updatecorrelator":"1"}
var mps = mps || {};
document.write('<scr'+'ipt id="mps-ext-load" src="//'+mpsopts.host+'/fetch/ext/load-'+mpscall.site+'.js"></scr'+'ipt>');
//--><!]]>
</script>
<script type="text/javascript" src="//www.nbcsports.com/sites/nbcsports.com/files/advagg_js/js__Y7yZleMJWCODBho1ln_hSuz9pe9N0ieTAYiDQrutTfM__LT8trz7iyEjtYadlF25AJiwubXOuQv_87zCVoyx8cU4__FnPHlXoj8IfbvL1keV4GMgRFJuL-j3m6YHoYNhvToAo.js"></script>
<script type="text/javascript">
<!--//--><![CDATA[//><!--
jQuery.extend(Drupal.settings, {"basePath":"\/",
"pathPrefix":"",
"ajaxPageState":{"theme":"nbcsports",
"theme_token":"Q78E44TPkEkz-VCJlo-kcNl4Lp2edgaq0tczu0Ow42o",
"jquery_version":"1.10",
"css":{"sites\/all\/modules\/contrib\/date\/date_popup\/themes\/datepicker.1.7.css":1,"sites\/all\/modules\/contrib\/video_filter\/video_filter.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/nbcsports-wrapper.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/layout.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/nbcsports-wrapper-overrides.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/custom-adaptive-layouts.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/vendor.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/blocks.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/colorbox_style.css":1,"sites\/all\/themes\/custom\/nbcsports\/css\/messages.css":1},"js":{"sites\/all\/modules\/contrib\/jquery_update\/replace\/jquery\/1.10\/jquery.min.js":1,"misc\/jquery.once.js":1,"misc\/drupal.js":1,"sites\/all\/modules\/contrib\/jquery_update\/replace\/ui\/ui\/minified\/jquery.ui.core.min.js":1,"sites\/all\/modules\/contrib\/jquery_update\/replace\/ui\/ui\/minified\/jquery.ui.widget.min.js":1,"sites\/all\/modules\/contrib\/jquery_update\/replace\/ui\/ui\/minified\/jquery.ui.tabs.min.js":1,"misc\/ajax.js":1,"sites\/all\/modules\/contrib\/jquery_update\/js\/jquery_update.js":1,"sites\/all\/modules\/contrib\/picture\/picturefill3\/picturefill.min.js":1,"sites\/all\/modules\/contrib\/picture\/picture.min.js":1,"sites\/all\/libraries\/colorbox\/jquery.colorbox-min.js":1,"sites\/all\/modules\/contrib\/colorbox\/js\/colorbox.js":1,"sites\/all\/modules\/contrib\/colorbox\/js\/colorbox_load.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/vendor\/owl.carousel.min.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/vendor\/fancySelect.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/vendor\/angular.min.js":1,"misc\/progress.js":1,"sites\/all\/modules\/custom\/nbcs_sitecatalyst\/js\/nbcs_s_wrapper.js":1,"sites\/all\/modules\/custom\/nbcs_sitecatalyst\/js\/nbcs_sitecatalyst.js":1,"sites\/all\/modules\/contrib\/colorbox_node\/colorbox_node.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/nbcsports-wrapper.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/slideout.min.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/google-webfont-loader.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/notify-fonts-ready.js":1,"sites\/all\/themes\/custom\/nbcsports\/js\/main.js":1,"sites\/all\/modules\/custom\/nbcs_live_extra\/js\/nbcs_live_extra_hub.js":1}},"colorbox":{"transition":"elastic","speed":"350","opacity":"0.85","slideshow":false,"slideshowAuto":true,"slideshowSpeed":"2500","slideshowStart":"start slideshow","slideshowStop":"stop slideshow","current":"{current} of {total}","previous":"\u00ab Prev","next":"Next \u00bb","close":"Close","overlayClose":true,"maxWidth":"100%","maxHeight":"100%","initialWidth":"300","initialHeight":"250","fixed":true,"scrolling":true,"mobiledetect":false,"mobiledevicewidth":"480px"},"vod_default_icon":"nbcsports-vod-default.jpg","liveExtra":{"events":[{"event_id":"16179","network":"NBCSN","event_title":"Swansea City v. Liverpool","sport":"Premier League","event_date_time":1475317800000,"event_date_time_end":1475330400000,"destination_url":"http:\/\/stream.nbcsports.com\/premier-league\/?pid=27125","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16205","network":"NBC Sports","event_title":"Ryder Cup Day 2","sport":"Golf","event_date_time":1475326800000,"event_date_time_end":1475364600000,"destination_url":"http:\/\/stream.golfchannel.com\/golf-generic?pid=27252","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16274","network":"Digital Only","event_title":"Premier League: Goal Rush","sport":"Soccer","event_date_time":1475328600000,"event_date_time_end":1475339400000,"destination_url":"http:\/\/stream.nbcsports.com\/nbc\/?pid=27084\u0026hls=1","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16180","network":"NBCSN","event_title":"Hull City v. Chelsea","sport":"Premier League","event_date_time":1475330400000,"event_date_time_end":1475339400000,"destination_url":"http:\/\/stream.nbcsports.com\/premier-league\/?pid=27121","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16183","network":"Digital Only","event_title":"West Ham v. Middlesbrough","sport":"Premier League","event_date_time":1475330400000,"event_date_time_end":1475339400000,"destination_url":"http:\/\/stream.nbcsports.com\/premier-league\/?pid=27119","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16184","network":"CNBC","event_title":"XFINITY Series: Dover Qualifying","sport":"NASCAR","event_date_time":1475335800000,"event_date_time_end":1475341200000,"destination_url":"http:\/\/stream.nbcsports.com\/nascar\/?pid=27214","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16209","network":"NBCSN","event_title":"Sprint Cup: Dover Final Practice","sport":"NASCAR","event_date_time":1475341200000,"event_date_time_end":1475346600000,"destination_url":"http:\/\/stream.nbcsports.com\/nascar\/?pid=27210","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16208","network":"Golf Live Extra","event_title":"Ryder Cup: The First Tee Experience","sport":"Golf","event_date_time":1475343000000,"event_date_time_end":1475344800000,"destination_url":"http:\/\/stream.golfchannel.com\/golf-generic?pid=27249","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16185","network":"NBCSN","event_title":"XFINITY Series: Dover Race","sport":"NASCAR","event_date_time":1475346600000,"event_date_time_end":1475361000000,"destination_url":"http:\/\/stream.nbcsports.com\/nascar\/?pid=27215","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16186","network":"NBCSN","event_title":"Breeders\u0027 Cup Challenge Series","sport":"Horse Racing","event_date_time":1475357400000,"event_date_time_end":1475366400000,"destination_url":"http:\/\/stream.nbcsports.com\/horse-racing\/?pid=27227","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16210","network":"Golf Channel","event_title":"Live from the Ryder Cup","sport":"Golf","event_date_time":1475362800000,"event_date_time_end":1475370000000,"destination_url":"http:\/\/stream.golfchannel.com\/golf-live-from?pid=27239","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16187","network":"NBCSN","event_title":"Gloucester v. Bath","sport":"Rugby","event_date_time":1475366400000,"event_date_time_end":1475373600000,"destination_url":"http:\/\/stream.nbcsports.com\/rugby\/?pid=27224","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16188","network":"NBCSN","event_title":"Premier League Match of the Day","sport":"Soccer","event_date_time":1475373600000,"event_date_time_end":1475380740000,"destination_url":"http:\/\/stream.nbcsports.com\/nbcsn\/?pid=15801","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16189","network":"Digital Only","event_title":"GP2 Series: Malaysia Race 2","sport":"F1","event_date_time":1475378400000,"event_date_time_end":1475381700000,"destination_url":"http:\/\/stream.nbcsports.com\/f1\/?pid=27221","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16190","network":"NBCSN","event_title":"Malaysia Grand Prix","sport":"F1","event_date_time":1475388000000,"event_date_time_end":1475402400000,"destination_url":"http:\/\/stream.nbcsports.com\/f1\/?pid=27222","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"},{"event_id":"16191","network":"NBCSN","event_title":"Man United v. Stoke City","sport":"Premier League","event_date_time":1475402400000,"event_date_time_end":1475414100000,"destination_url":"http:\/\/stream.nbcsports.com\/premier-league\/?pid=27123","dma_list":"","field_image_fid":null,"event_type":"0","feature_image_url":"\/sites\/all\/modules\/custom\/nbcs_live_extra\/images\/default_live_extra_100x56.jpg"}],"eventsDay":null},"colorbox_node":{"width":"1260px","height":"650px"},"nbcs":{"site-group":"nbcs","site":"nbcsports.com"}});
//--><!]]>
</script>
</head>
<body class="html not-front not-logged-in page-live nbcs-site has-one-sidebar has-sidebar-second sidebar-second" >

  <header id="nbcsports-fixed-header">

    <!-- start Primary Nav -->
    <nav id="nbcsports-primary-nav">
                  <a href="http://www.nbcsports.com" id="home-page-link"></a>  <a href="http://www.nbcsports.com/nfl" class="dynamic-nav-link" data-tid="6">NFL</a>  <a href="http://www.nbcsports.com/nhl" class="dynamic-nav-link" data-tid="11">NHL</a>  <a href="http://www.nbcsports.com/nba" class="dynamic-nav-link" data-tid="16">NBA</a>  <a href="http://www.nbcsports.com/mlb" class="dynamic-nav-link" data-tid="21">MLB</a>  <a href="http://www.nbcsports.com/soccer" class="dynamic-nav-link" data-tid="26">Soccer</a>  <a href="http://www.nbcsports.com/nascar" class="dynamic-nav-link" data-tid="46">NASCAR</a>  <a href="http://www.golfchannel.com/" class="dynamic-nav-link">Golf</a>  <a href="http://www.nbcsports.com/motors" class="dynamic-nav-link" data-tid="41">Motors</a>  <a href="http://www.nbcsports.com/college-football" class="dynamic-nav-link" data-tid="31">NCAA FB</a>  <a href="http://www.nbcsports.com/olympic-sports" class="dynamic-nav-link">Oly</a>  <a href="http://www.nbcsports.com/horse-racing" class="dynamic-nav-link" data-tid="61">Horses</a>  <a href="http://www.nbcsports.com/rugby" class="dynamic-nav-link" data-tid="101">Rugby</a>  <a href="http://www.nbcsports.com/cycling" class="dynamic-nav-link" data-tid="56">CYC</a>  <a href="http://www.nbcsports.com/tennis" class="dynamic-nav-link" data-tid="51">Tennis</a>  <a href="http://www.nbcsports.com/college-basketball" class="dynamic-nav-link" data-tid="36">NCAA BK</a>  <a href="http://www.nbcsports.com/boxing" class="dynamic-nav-link" data-tid="71">Boxing</a>  <a href="http://www.nbcsports.com/mma-0" class="dynamic-nav-link" data-tid="76">MMA</a>  <a href="http://www.nbcsports.com/national-dog-show" class="dynamic-nav-link">Dog Show</a>  <a href="http://www.dewtour.com/" class="dynamic-nav-link">Dew Tour</a>  <a href="http://www.redbullsignatureseries.com/" class="dynamic-nav-link">Red Bull</a>  <a href="http://www.nbcsports.com/spartan-race" class="dynamic-nav-link">Spartan Race</a>  <a href="http://www.nbcsnauto360.com/" class="dynamic-nav-link">Auto programming</a>  <a href="http://www.nbcsnoutdoors360.com/" class="dynamic-nav-link">Outdoors</a>  <a href="http://shop.nbcsports.com/" class="dynamic-nav-link">Shop</a>  <a href="http://www.nbcsports.com/gold" class="dynamic-nav-link">NBC Sports Gold</a>
<span id="more-nav-btn" class="no-select no-more-links" unselectable="on" tabindex="0">MORE</span>
    </nav>
    <!-- end Primary Nav -->

    <!-- start More Nav -->
    <nav id="more-nav" class="no-more-links">
    </nav>
    <!-- end More Nav -->

    <!-- start Secondary Nav -->
    <nav id="nbcsports-secondary-nav">

<div class="nbcsports-secondary-nav__wrapper">

        <a href="http://www.nbcsports.com/live">Live and Upcoming</a>        <a href="http://www.nbcsports.com/video">Clips and Highlights</a>        <a href="http://www.nbcsports.com/full-event-replays">Full Event Replays</a>        <a href="http://www.nbcsports.com/sports-mobile">Apps</a>        <a href="http://www.nbcsports.com/gold" class="gold-nav">Gold</a>        <a href="http://www.golfchannel.com/" class="golf-channel-nav">Golf Channel</a>        <a href="http://www.rotoworld.com/" class="rotoworld-nav">Rotoworld</a>        <a href="http://www.comcastsportsnet.com/" class="rsn-local-nav">Comcast SportsNet</a>
  <div id="nbcsports-search-box">
    <form action="http://www.nbcsports.com/search" class="search-form" method="get" autocomplete="off">
      <span class="icon"></span>
      <label>Search</label>
      <input type="search" name="search_api_views_fulltext" id="nbcsports-search" placeholder="Search NBC Sports" tabindex="0"/>
    </form>
  </div>

</div>
    </nav>
    <!-- end Secondary Nav -->

  </header>
  <!-- end #nbcsports-fixed-header -->

  <div id="wrapper">
    <div id="outer">
      <!-- Start leaderboard -->
      <div id="nbcsports-leaderboard">
                    <div class="region region-top-region">
    <div id="block-mps-topbanner" class="block block-mps">


  <div class="content">
    <div id="topbanner"><script>typeof(mps.getAd)=='function' && document.write(mps.getAd('topbanner'));</script></div>  </div>
</div>
  </div>
              </div>
      <!-- End leaderboard -->

      <!-- Start scoreboard -->
              <div id="nbcsports-scoreboard">
            <div class="region region-scoreboard">
    <div class="scoreboard with_ad">
  <iframe id="scoreboard-ticker" src="http://scores.nbcsports.com/score_ticker/index.html?sports=CFB%7CEPL%7CNFL%7CMLB%7CNASCAR%7CGOLF%7CFORM1%7CMLS"></iframe>
</div>
<div id="block-mps-scorelogo" class="block block-mps scoreboard-ad">


  <div class="content">
    <div id="scorelogo"><script>typeof(mps.getAd)=='function' && document.write(mps.getAd('scorelogo'));</script></div>  </div>
</div>
  </div>
        </div>
            <!-- Start scoreboard -->

      <!-- Start content wrapper -->
      <div id="nbcsports-content-wrapper">

        <aside id="nbcsports-live-extra-widget">
                  </aside>

        <!-- Start main content -->
        <div id="nbcsports-main" class="clearfix">

                      <div class="tabs">
                          </div>



          <div class="main-layout">
            <div class="layout_full-minus-right layout_first dynamic-lede">
              <div class="content-inner">
                                              </div>
            </div>

            <div class="layout_full-minus-right layout_second dynamic-lede">
                <div class="content-inner">
                    <div class="region region-content">
    <div id="block-system-main" class="block block-system">


  <div class="content">
    <div class="live-extra-page"  data-ng-app="LiveExtra" data-ng-controller="LiveExtraMainController">
  <div class="tabs_custom live-extra-page__tabs">
    <div class="tabs_custom__mobile-switcher desktop-large_hidden desktop-small_hidden tablet_hidden">
      <span class="tabs_custom__mobile-current">Current</span>
    </div>
    <ul class="tabs_custom__nav mobile_hidden"><li>
    <a class="anchor tabs_custom__anchor anchor_live-upcoming" href="#live-upcoming" data-section="live">
      Live &amp; Upcoming
    </a>
  </li>
<li>
    <a class="anchor tabs_custom__anchor anchor_full-events-replay" href="#full-events-replays" data-section="live">
      Full Event Replays
    </a>
  </li>
<li><a href="/video" class="anchor tabs_custom__anchor">Clips and Highlights</a></li>
<li><a href="/gold" class="anchor tabs_custom__anchor">Gold</a></li>
</ul>    <div class="live-upcoming-tab" id="live-upcoming">
      <div class="live-upcoming-filter">
        <div class="form__inline-group">
                      <form class="live-upcoming-filter__form form form_inline" action="/live" method="get">
              <div class="form__item select-box select-box_custom live-upcoming-filter__sports">
                <select class="select select-box__select">
                  <option value="All">All Sports</option>
                                      <option value="Boxing">Boxing</option>
                                      <option value="F1">F1</option>
                                      <option value="Fantasy Football">Fantasy Football</option>
                                      <option value="General">General</option>
                                      <option value="Golf">Golf</option>
                                      <option value="Horse Racing">Horse Racing</option>
                                      <option value="IndyCar">IndyCar</option>
                                      <option value="NASCAR">NASCAR</option>
                                      <option value="NCAA Football">NCAA Football</option>
                                      <option value="NFL">NFL</option>
                                      <option value="NHL">NHL</option>
                                      <option value="Notre Dame">Notre Dame</option>
                                      <option value="Olympic Sports">Olympic Sports</option>
                                      <option value="Premier League">Premier League</option>
                                      <option value="Rugby">Rugby</option>
                                      <option value="Soccer">Soccer</option>
                                      <option value="WSOF">WSOF</option>
                                  </select>
              </div>
            </form>
                    <form class="live-upcoming-filter__form form form_inline mobile_hidden" action="/tv-listings" method="get">
            <div class="form__item input-box live-upcoming-filter__zip">
              <input type="hidden" name="ap" value="cf" />
              <input type="hidden" name="v" value="1" />
              <input type="hidden" name="aid" value="nbcsports" />
              <input type="hidden" name="stnlt" value="15952,48639" />
              <!-- for more than 2 stations, separate with commas ?????,????? -->
              <input type="hidden" name="nstnlt" value="false" />
              <input type="hidden" name="rty" value="json" />

              <input type="text" class="input form__input" name="zip" placeholder="Enter Zipcode" id="channel-finder-zip" />
            </div>
            <div class="form__item button-box live-upcoming-filter__submit">
              <input type="submit" class="button form__submit" value="Find Channel" />
            </div>
          </form>
        </div>
      </div>
      <div class="live-upcoming-time-line">
        <div class="time-line owl-carousel mobile-hidden">
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-01" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/01</span>
                <span class="date-day">saturday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-02" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/02</span>
                <span class="date-day">sunday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-03" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/03</span>
                <span class="date-day">monday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-04" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/04</span>
                <span class="date-day">tuesday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-05" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/05</span>
                <span class="date-day">wednesday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-06" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/06</span>
                <span class="date-day">thursday</span>
              </span>
              </a>
            </div>
                      <div class="time-line__item">
              <a class="link" href="/live/2016-10-07" data-ng-click="fetchEvents($event); $event.preventDefault(); $event.stopPropagation();">
              <span class="time-lime__date">
                <span class="date-digit">10/07</span>
                <span class="date-day">friday</span>
              </span>
              </a>
            </div>
                  </div>
      </div>
      <div id="live-extra-hub-live-upcoming-wrapper" data-ng-cloak>
        <div class="live-upcoming-list live-upcomming-list">
          <div class="live-upcoming-list__loading" data-ng-class="{ 'loading': loading }"></div>
          <div class="live-upcoming-list__events">
            <div class="events-list live-upcoming-list__events-list live-upcoming-list__events-list_type-upcoming">
              <div class="events-list__list-wrapper events-list__list-wrapper_live" data-ng-show="events.live.length">
                <div class="title events-list__title events-list__title_type-live">
                  <span class="button button_live">Live</span>
                </div>
                <ul class="list events-list__list">
                  <li class="list__item event-active" data-ng-repeat="event in events.live" data-live-extra-event></li>
                </ul>
              </div>
              <div class="events-list__list-wrapper events-list__list-wrapper_upcoming" data-ng-show="!isUpcomingEmpty">
                <div class="title events-list__title events-list__title_type-upcoming">
                  <span class="button button_upcoming">Upcoming</span>
                </div>
                <div class="events-list__subgroup-wrapper" data-ng-repeat="(timestamp, subgroup) in events.upcoming">
                  <div class="sub-title events-list__sub-title events-list__sub-title_type-upcoming" data-subgroup="{{ timestamp }}">
                    {{ timestamp | formatAMPM }}
                  </div>
                  <ul class="list events-list__list">
                    <li class="list__item" data-ng-repeat="event in subgroup" data-live-extra-event data-ng-class="{ 'event-active' : event.isActive }"</li>
                  </ul>
                </div>
              </div>
              <div id="live-extra-hub-live-upcoming__no_events" data-ng-show="isUpcomingEmpty">{{ noEventsMessage }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="full-events-replay-tab" id="full-events-replays">
  <div class="videos-list-filter">
    <form class="videos-list-filter__form form form_inline" action="/live#full-events-replays" method="get">
      <div class="form__inline-group">
                  <div class="form__item select-box select-box_custom videos-list-filter__sports">
            <select class="select select-box__select">
              <option value="All">All Sports</option>
                              <option value="Boxing">Boxing</option>
                              <option value="F1">F1</option>
                              <option value="Fantasy Football">Fantasy Football</option>
                              <option value="General">General</option>
                              <option value="Golf">Golf</option>
                              <option value="Horse Racing">Horse Racing</option>
                              <option value="IndyCar">IndyCar</option>
                              <option value="NASCAR">NASCAR</option>
                              <option value="NCAA Football">NCAA Football</option>
                              <option value="NFL">NFL</option>
                              <option value="NHL">NHL</option>
                              <option value="Notre Dame">Notre Dame</option>
                              <option value="Olympic Sports">Olympic Sports</option>
                              <option value="Premier League">Premier League</option>
                              <option value="Rugby">Rugby</option>
                              <option value="Soccer">Soccer</option>
                              <option value="WSOF">WSOF</option>
                          </select>
          </div>
                <div class="form__item button-box videos-list-filter__submit">
          <input type="submit" class="button" value="Find" />
        </div>
      </div>
    </form>
  </div>
  <div class="live-extra-page__videos">
    <div id="live-extra-hub-replays__no_events">No replays found.</div>
                  <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NASCAR">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nascar/?pid=27209">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sprint Cup: Dover Second Practice" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sprint Cup: Dover Second Practice</span>
                  <span class="video-event-thumb__event-type">NASCAR</span>
                  <span class="video-event-thumb__event-day">10/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Premier League">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/premier-league/?pid=27118">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Watford v. Bournemouth" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Watford v. Bournemouth</span>
                  <span class="video-event-thumb__event-type">Premier League</span>
                  <span class="video-event-thumb__event-day">10/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Premier League">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/premier-league/?pid=27124">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sunderland v. West Brom" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sunderland v. West Brom</span>
                  <span class="video-event-thumb__event-type">Premier League</span>
                  <span class="video-event-thumb__event-day">10/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Premier League">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/premier-league/?pid=27119">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/10/01/27119.jpg?itok=0SgSWmrV" alt="West Ham v Middlesbrough" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">West Ham v Middlesbrough</span>
                  <span class="video-event-thumb__event-type">Premier League</span>
                  <span class="video-event-thumb__event-day">10/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27220">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Malaysia Grand Prix Qualifying" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Malaysia Grand Prix Qualifying</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">10/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27218">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="GP2 Series: Malaysia Race 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">GP2 Series: Malaysia Race 1</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NCAA Football">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbcsn/generic?pid=27229">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Penn vs. Dartmouth" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Penn vs. Dartmouth</span>
                  <span class="video-event-thumb__event-type">NCAA Football</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://www.nbcsports.com/video/premier-league-news-sept-30-2016">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Premier League News: Sept. 30, 2016" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier League News: Sept. 30, 2016</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/rugby/?pid=27223">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Northampton Saints v. Exeter Chiefs" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Northampton Saints v. Exeter Chiefs</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NASCAR">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nascar/?pid=27207">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sprint Cup: Dover Practice" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sprint Cup: Dover Practice</span>
                  <span class="video-event-thumb__event-type">NASCAR</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27217">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Malaysia Grand Prix Practice 2" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Malaysia Grand Prix Practice 2</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/30</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27216">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Malaysia Grand Prix Practice 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Malaysia Grand Prix Practice 1</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/29</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://www.nbcsports.com/video/premier-league-news-sept-29-2016">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Premier League News: Sept. 29, 2016" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier League News: Sept. 29, 2016</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/29</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Fantasy Football">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/fantasy-football-live?pid=27484">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Rotoworld Fantasy Football Kickoff" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Rotoworld Fantasy Football Kickoff</span>
                  <span class="video-event-thumb__event-type">Fantasy Football</span>
                  <span class="video-event-thumb__event-day">09/29</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://www.nbcsports.com/video/premier-league-news-sept-28-2016">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Premier League News: Sept. 28, 2016" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier League News: Sept. 28, 2016</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/28</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbcsn/index_nbcsn-generic.html?pid=27320">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Spartan Race: Asheville" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Spartan Race: Asheville</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/27</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://www.nbcsports.com/video/premier-league-news-sept-27-2016">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Premier League News: Sept. 27, 2016" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier League News: Sept. 27, 2016</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/27</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27322">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Men in Blazers" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Men in Blazers</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27326">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="In The Zone Ep. 16.17" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">In The Zone Ep. 16.17</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27325">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="FEI Equestrian World Ep. 16.8" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">FEI Equestrian World Ep. 16.8</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27324">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="WTS Cozumel Magazine" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">WTS Cozumel Magazine</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://www.nbcsports.com/video/premier-league-news-september-26-2016">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Premier League News: September 26, 2016" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier League News: September 26, 2016</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27323">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="FNIA Aquatics Magazine (Episode 9)" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">FNIA Aquatics Magazine (Episode 9)</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=21727">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Rugby World Magazine (Episode 34)" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Rugby World Magazine (Episode 34)</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/26</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/rugby/?pid=27102">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Leicester Tigers vs. Bath" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Leicester Tigers vs. Bath</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/25</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NASCAR">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nascar-sprint/?pid=27272">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sprint Cup: New Hampshire Race" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sprint Cup: New Hampshire Race</span>
                  <span class="video-event-thumb__event-type">NASCAR</span>
                  <span class="video-event-thumb__event-day">09/25</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27090">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Berlin Marathon" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Berlin Marathon</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/25</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27194">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Cyclocross World Cup: Men's Elite Iowa" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Cyclocross World Cup: Men's Elite Iowa</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/24</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27089">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Cyclocross World Cup: Women's Elite Iowa" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Cyclocross World Cup: Women's Elite Iowa</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/24</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Notre Dame">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/notre-dame/?pid=27105">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Duke vs. Notre Dame" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Duke vs. Notre Dame</span>
                  <span class="video-event-thumb__event-type">Notre Dame</span>
                  <span class="video-event-thumb__event-day">09/24</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/rugby/?pid=27101">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Northampton vs. Wasps" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Northampton vs. Wasps</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/24</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27182">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Cyclo-Cross World Cup: Women's Elite Las Vegas" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Cyclo-Cross World Cup: Women's Elite Las Vegas</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/21</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27088&hls=1">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Cyclo-Cross World Cup: Men's Elite Las Vegas" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Cyclo-Cross World Cup: Men's Elite Las Vegas</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/21</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27108">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Minor League Baseball Championships" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Minor League Baseball Championships</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/20</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27163">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Rugby World Episode 33" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Rugby World Episode 33</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/19</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="IndyCar">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/indy/?pid=27066">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Grand Prix of Sonoma" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Grand Prix of Sonoma</span>
                  <span class="video-event-thumb__event-type">IndyCar</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27073">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="ITU World Triathlon (Men)" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">ITU World Triathlon (Men)</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27071">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="FIVB Beach Volleyball World Tour: Men's Final" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">FIVB Beach Volleyball World Tour: Men's Final</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27072">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="2016 Rio Paralympic Games" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">2016 Rio Paralympic Games</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27150">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="FIVB Beach Volleyball Gold Match" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">FIVB Beach Volleyball Gold Match</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27068">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Global Rallycross: Seattle" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Global Rallycross: Seattle</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27149">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="FIVB Beach Volleyball World Tour: Women's Final" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">FIVB Beach Volleyball World Tour: Women's Final</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/rugby/?pid=27076">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Leicester Tigers vs. Newcastle Falcons" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Leicester Tigers vs. Newcastle Falcons</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27049">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Singapore Grand Prix" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Singapore Grand Prix</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/18</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Notre Dame">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/notre-dame/?pid=27067">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Michigan State vs. Notre Dame" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Michigan State vs. Notre Dame</span>
                  <span class="video-event-thumb__event-type">Notre Dame</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="IndyCar">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/indy/?pid=27065">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="IndyCar Qualifying: Grand Prix of Sonoma" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">IndyCar Qualifying: Grand Prix of Sonoma</span>
                  <span class="video-event-thumb__event-type">IndyCar</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27070">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="ITU World Triathlon (Women)" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">ITU World Triathlon (Women)</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="IndyCar">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/indy/?pid=27064">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sonoma Practice" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sonoma Practice</span>
                  <span class="video-event-thumb__event-type">IndyCar</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27069">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Swatch FIVB World Tour Women Finals" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Swatch FIVB World Tour Women Finals</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NASCAR">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nascar/?pid=27053">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Sprint Cup: Chicagoland Practice" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Sprint Cup: Chicagoland Practice</span>
                  <span class="video-event-thumb__event-type">NASCAR</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27048">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Singapore Grand Prix Qualifying" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Singapore Grand Prix Qualifying</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27047">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Singapore Grand Prix Practice 3" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Singapore Grand Prix Practice 3</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/17</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=19316">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="The Men in Blazers Show" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">The Men in Blazers Show</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/16</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=27045">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/27045.jpg?itok=sl0WkyHw" alt="Singapore Grand Prix Practice 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Singapore Grand Prix Practice 1</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/16</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Soccer">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27016">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/mib_1.jpg?itok=XHmL9TWB" alt="Men in Blazers" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Men in Blazers</span>
                  <span class="video-event-thumb__event-type">Soccer</span>
                  <span class="video-event-thumb__event-day">09/12</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27002">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/inthezone_1.jpg?itok=JSR-dyeS" alt="In the Zone, Episode 16.15" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">In the Zone, Episode 16.15</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/12</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=27001">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/triathlon_1.jpg?itok=EcIfImKr" alt="ITU World Triathlon Series: Edmonton" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">ITU World Triathlon Series: Edmonton</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/12</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26669">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26669_1920x1080.jpg?itok=bCFvmHfj" alt="Vuelta a Espana: Stage 21" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a Espana: Stage 21</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/11</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26798">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26798.jpg?itok=0FZDVVPS" alt="Tour of Britain: Stage 8" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 8</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/11</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbcsn/index_nbcsn-generic.html?pid=27030">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26997.jpg?itok=WrA4Grs-" alt="Vuelta a España: Madrid Challenge" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Madrid Challenge</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/11</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=26853">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/americascup.jpg?itok=XCCWA5w-" alt="America's Cup World Series: Toulon" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">America's Cup World Series: Toulon</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/11</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26800">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26799_1.jpg?itok=cUBKno1g" alt="UCI Downhill Championships" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">UCI Downhill Championships</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/11</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Rugby">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/rugby/?pid=26867">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/cr_z8xbvyaagvfb.jpg?itok=_Es_O57v" alt="Leicester City vs. Wasps" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Leicester City vs. Wasps</span>
                  <span class="video-event-thumb__event-type">Rugby</span>
                  <span class="video-event-thumb__event-day">09/10</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26797">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26797_0.jpg?itok=DC28T2DD" alt="Tour of Britain: Stage 7" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 7</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/10</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26668">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26668_1920x1080.jpg?itok=3_B0EUqC" alt="Vuelta a España: Stage 20" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 20</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/10</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Notre Dame">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/notre-dame/?pid=26849">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/nd_0.jpg?itok=ZA12LTI4" alt="Nevada vs. Notre Dame" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Nevada vs. Notre Dame</span>
                  <span class="video-event-thumb__event-type">Notre Dame</span>
                  <span class="video-event-thumb__event-day">09/10</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26799">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26799.jpg?itok=Ia6ymGLr" alt="Four-Cross Championships" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Four-Cross Championships</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/09</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26796">
                      <img class="image" src="/sites/all/modules/custom/nbcs_live_extra/images/default_live_extra_170x96.jpg" alt="Tour of Britain: Stage 6" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 6</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/09</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Fantasy Football">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=26994">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26994.jpg?itok=IIR378Jo" alt="Rotoworld Fantasy Football Kickoff" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Rotoworld Fantasy Football Kickoff</span>
                  <span class="video-event-thumb__event-type">Fantasy Football</span>
                  <span class="video-event-thumb__event-day">09/08</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26666">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26666_1920x1080.jpg?itok=XNmG3thb" alt="Vuelta a España: Stage 18" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 18</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/08</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26795">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26795_0.jpg?itok=j5WDL_QK" alt="Tour of Britain: Stage 5" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 5</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/08</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26665">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26665_1920x1080.jpg?itok=n_z2DVam" alt="Vuelta a España: Stage 17" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 17</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/07</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26794">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26794_0.jpg?itok=PBKg-QA8" alt="Tour of Britain: Stage 4" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 4</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/07</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26793">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26793_0.jpg?itok=qrGGP15u" alt="Tour of Britain: Stage 3" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 3</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/06</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26664">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26664_1920x1080.jpg?itok=PoPlVAsh" alt="Vuelta a España: Stage 16" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 16</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/05</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26792">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26792_0.jpg?itok=kDZZktOo" alt="Tour of Britain: Stage 2" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 2</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/05</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="NASCAR">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/indy/?pid=26813">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26813_0.jpg?itok=0wOEdSSM" alt="Watkins Glen Race" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Watkins Glen Race</span>
                  <span class="video-event-thumb__event-type">NASCAR</span>
                  <span class="video-event-thumb__event-day">09/04</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Olympic Sports">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/nbc/?pid=26841">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26841_0.jpg?itok=TsKZKjm0" alt="AVP Volleyball: Chicago" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">AVP Volleyball: Chicago</span>
                  <span class="video-event-thumb__event-type">Olympic Sports</span>
                  <span class="video-event-thumb__event-day">09/04</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26663">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26663_1920x1080.jpg?itok=Y-k-tMyR" alt="Vuelta a Espana: Stage 15" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a Espana: Stage 15</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/04</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26791">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26791_0.jpg?itok=orgXlE3k" alt="Tour of Britain: Stage 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour of Britain: Stage 1</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/04</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="F1">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/f1/?pid=26807">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/gp2.jpg?itok=rpapdJWh" alt="GP2 Series: Italy Race 2" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">GP2 Series: Italy Race 2</span>
                  <span class="video-event-thumb__event-type">F1</span>
                  <span class="video-event-thumb__event-day">09/04</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Boxing">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/boxing/?pid=26844">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/pbc_0_1.jpg?itok=Gc39XPbd" alt="Premier Boxing Champions" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Premier Boxing Champions</span>
                  <span class="video-event-thumb__event-type">Boxing</span>
                  <span class="video-event-thumb__event-day">09/03</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="Horse Racing">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/horse-racing/?pid=26843">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26843_0.jpg?itok=iKCK7wz4" alt="Breeders' Cup Challenge Series: Woodward Stakes" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Breeders' Cup Challenge Series: Woodward Stakes</span>
                  <span class="video-event-thumb__event-type">Horse Racing</span>
                  <span class="video-event-thumb__event-day">09/03</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26662">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26662_1920x1080.jpg?itok=ECwgl6WD" alt="Vuelta a Espana: Stage 14" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a Espana: Stage 14</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/03</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="IndyCar">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/indy/?pid=26810">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/indycar.jpg?itok=0tanTXjH" alt="Watkins Glen Practice" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Watkins Glen Practice</span>
                  <span class="video-event-thumb__event-type">IndyCar</span>
                  <span class="video-event-thumb__event-day">09/02</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26066">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26661_1920x1080.jpg?itok=iFmDjyMA" alt="Vuelta a España: Stage 13" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 13</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/02</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26660">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26660_1920x1080.jpg?itok=3mhje7Qc" alt="Vuelta a España: Stage 12" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 12</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">09/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26659">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26659_1920x1080.jpg?itok=62H9ev_y" alt="Vuelta a España: Stage 11" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 11</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/31</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26658">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26658_1920x1080.jpg?itok=HkDQlgg6" alt="Vuelta a España: Stage 10" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 10</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/29</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26657">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26657_1920x1080.jpg?itok=Uc3YmgrP" alt="Vuelta a España: Stage 9" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 9</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/28</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26656">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26656_1920x1080.jpg?itok=jbgkxuK1" alt="Vuelta a España: Stage 8" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 8</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/27</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26655">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26655_1920x1080.jpg?itok=jd8GHLCF" alt="Vuelta a España: Stage 7" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 7</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/26</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26654">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26654_1920x1080.jpg?itok=WbnKzY53" alt="Vuelta a España: Stage 6" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 6</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/25</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26653">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26653_1920x1080.jpg?itok=GiQXl9lA" alt="Vuelta a España: Stage 5" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 5</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/24</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26652">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26652_1920x1080.jpg?itok=PiwC2OAr" alt="Vuelta a España: Stage 4" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 4</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/23</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26650">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26650_1920x1080.jpg?itok=OMs3Zg1I" alt="Vuelta a España: Stage 2" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 2</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/21</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/gold-cycling/?pid=26649">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/26649_1920x1080.jpg?itok=IskYfC-c" alt="Vuelta a España: Stage 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Vuelta a España: Stage 1</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">08/20</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22508">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22508.jpg?itok=3bqWa1U5" alt="Tour de France: Stage 21" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 21</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22507">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22507.jpg?itok=u0ryHzQf" alt="Tour de France: Stage 20" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 20</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22506">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22506.jpg?itok=FQz_JJX3" alt="Tour de France: Stage 19" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 19</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22505">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22505.jpg?itok=JwrVrsoB" alt="Tour de France: Stage 18" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 18</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22504">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22504.jpg?itok=AP6Z3AbF" alt="Tour de France: Stage 17" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 17</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22503">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22503.jpg?itok=S4fg3Tkg" alt="Tour de France: Stage 16" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 16</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22502">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22502.jpg?itok=9OtWQwiL" alt="Tour de France: Stage 15" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 15</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22501">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22501.jpg?itok=sqJFgAGL" alt="Tour de France: Stage 14" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 14</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22500">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22500.jpg?itok=q4bvCXHy" alt="Tour de France: Stage 13" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 13</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22499">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22499.jpg?itok=wDZlq7_4" alt="Tour de France: Stage 12" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 12</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22498">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22498.jpg?itok=785xoTOZ" alt="Tour de France: Stage 11" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 11</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22497">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22497.jpg?itok=yFfRGFRT" alt="Tour de France: Stage 10" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 10</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22496">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22496.jpg?itok=dWlmRRLs" alt="Tour de France: Stage 9" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 9</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22495">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22495.jpg?itok=On5POktQ" alt="Tour de France: Stage 8" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 8</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22494">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22494.jpg?itok=BK0dOb1O" alt="Tour de France: Stage 7" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 7</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22493">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22493.jpg?itok=i-_9Jldk" alt="Tour de France: Stage 6" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 6</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22308">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22308.jpg?itok=JaGSWUme" alt="Tour de France: Stage 5" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 5</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22307">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22307.jpg?itok=rMygADXj" alt="Tour de France: Stage 4" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 4</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22306">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22306.jpg?itok=ANFLAN2w" alt="Tour de France: Stage 3" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 3</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
              <div class="live-extra-page__videos-row">
          <ul class="video-events-list">
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=27029">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22305.jpg?itok=0GHtE6LZ" alt="Tour de France: Stage 2" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 2</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                          <li class="video-events-list__item">
                <div class="video-event-thumb" data-sport="General">
                  <div class="video-event-thumb__image">
                    <a class="link link_type-block link_type-play" href="http://stream.nbcsports.com/tour-de-france/?pid=22304">
                      <img class="image" src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/scale_and_crop_168x95/public/2016/09/16/22304.jpg?itok=E0tGurN0" alt="Tour de France: Stage 1" />
                      <span class="play-btn"></span>
                    </a>
                  </div>
                  <span class="video-event-thumb__event-name">Tour de France: Stage 1</span>
                  <span class="video-event-thumb__event-type">General</span>
                  <span class="video-event-thumb__event-day">07/01</span>
                </div>
              </li>
                      </ul>
        </div>
            </div>

</div>
  </div>
</div>
  </div>
</div>
  </div>
                </div>
            </div>

                        <aside id="nbcsports-right-rail" class="sidebar sidebar_second">
                <div class="region region-right-rail">
    <div id="block-bean-live-extra-landing-page-spotligh" class="block block-bean">


  <div class="content">

<div class="featured-videos-list featured-videos-list_carousel">
      <div  class="featured-videos-list__item first">
  <div class="featured-videos-list__thumb">
          <div class="featured-videos-list__image">
                  <a class="link link_type-play link_type-block"
             href="/video/arsene-wenger-reflects-highs-and-lows-20-years-arsenal">
            <picture >
<!--[if IE 9]><video style="display: none;"><![endif]-->
<source srcset="http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=C7e2esDi&amp;timestamp=1475348582 1x, http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_2x_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=tjaymNKC&amp;timestamp=1475348582 2x" media="(min-width: 1260px)" />
<source srcset="http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=C7e2esDi&amp;timestamp=1475348582 1x, http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_2x_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=tjaymNKC&amp;timestamp=1475348582 2x" media="(min-width: 1024px)" />
<source srcset="http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=C7e2esDi&amp;timestamp=1475348582 1x, http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_2x_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=tjaymNKC&amp;timestamp=1475348582 2x" media="(min-width: 768px)" />
<source srcset="http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=C7e2esDi&amp;timestamp=1475348582 1x, http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_2x_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=tjaymNKC&amp;timestamp=1475348582 2x" media="(min-width: 320px)" />
<!--[if IE 9]></video><![endif]-->
<img  src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/double_2__desktop_large_/public/2016/10/01/nbc_epl_wengerdownloadfull_161001.jpg?itok=C7e2esDi&amp;timestamp=1475348582" alt="" title="" />
</picture>            <span class="play-btn"></span>
                          <div class="video-duration-overlay">25:05</div>
                      </a>
              </div>
        <div class="featured-videos-list__desc">
      <span class="featured-videos-list__event-title">
        <a href="/video/arsene-wenger-reflects-highs-and-lows-20-years-arsenal">Wenger on 20 years at Arsenal</a>
      </span>
              <span class="featured-videos-list__event-description">
          In the latest PL Download, Roger Bennett talks to Arsene Wenger about his
20-year reign at Arsenal.
        </span>
          </div>
  </div>
</div>
      <div  class="featured-videos-list__item">
  <div class="featured-videos-list__thumb">
          <div class="featured-videos-list__image">
                  <a class="link link_type-play link_type-block"
             href="/video/tony-stewart-joins-nascar-america-reflect-hall-fame-racing-career">
            <img src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/half_1__desktop_large_extra_live_/public/2016/09/28/nbc_nas_tonystewart_160928.jpg?itok=FGYU7RAT" width="220" height="123" alt="" />            <span class="play-btn"></span>
                          <div class="video-duration-overlay">4:38</div>
                      </a>
              </div>
        <div class="featured-videos-list__desc">
      <span class="featured-videos-list__event-title">
        <a href="/video/tony-stewart-joins-nascar-america-reflect-hall-fame-racing-career">Tony Stewart reflects on career</a>
      </span>
          </div>
  </div>
</div>
      <div  class="featured-videos-list__item">
  <div class="featured-videos-list__thumb">
          <div class="featured-videos-list__image">
                  <a class="link link_type-play link_type-block"
             href="/video/chelsea-end-slide-2-0-win-v-hull">
            <img src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/half_1__desktop_large_extra_live_/public/2016/10/01/nbc_epl_stamfordhighlightschehul_161001.jpg?itok=-Xc1upBw" width="220" height="123" alt="" />            <span class="play-btn"></span>
                          <div class="video-duration-overlay">3:32</div>
                      </a>
              </div>
        <div class="featured-videos-list__desc">
      <span class="featured-videos-list__event-title">
        <a href="/video/chelsea-end-slide-2-0-win-v-hull">Chelsea end slide in win v. Hull</a>
      </span>
          </div>
  </div>
</div>
      <div  class="featured-videos-list__item">
  <div class="featured-videos-list__thumb">
          <div class="featured-videos-list__image">
                  <a class="link link_type-play link_type-block"
             href="/video/watford-and-bournemouth-finish-2-2-thriller-vicarage-road">
            <img src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/half_1__desktop_large_extra_live_/public/2016/10/01/2016-10-01t15-57-16.9z-1280x720.jpg?itok=lGSMhVJi" width="220" height="123" alt="" />            <span class="play-btn"></span>
                          <div class="video-duration-overlay">2:28</div>
                      </a>
              </div>
        <div class="featured-videos-list__desc">
      <span class="featured-videos-list__event-title">
        <a href="/video/watford-and-bournemouth-finish-2-2-thriller-vicarage-road">Watford, Bournemouth play to thrilling draw</a>
      </span>
          </div>
  </div>
</div>
      <div  class="featured-videos-list__item">
  <div class="featured-videos-list__thumb">
          <div class="featured-videos-list__image">
                  <a class="link link_type-play link_type-block"
             href="/video/will-josh-gordon-ever-play-cleveland-browns-nfl">
            <img src="http://www.nbcsports.com/sites/nbcsports.com/files/styles/half_1__desktop_large_extra_live_/public/2016/09/30/nbc_nfl_pft_gordonrehabv2_160930.jpg?itok=mmSrS2P7" width="220" height="123" alt="" />            <span class="play-btn"></span>
                          <div class="video-duration-overlay">1:02</div>
                      </a>
              </div>
        <div class="featured-videos-list__desc">
      <span class="featured-videos-list__event-title">
        <a href="/video/will-josh-gordon-ever-play-cleveland-browns-nfl">Will Josh Gordon ever play in NFL?</a>
      </span>
          </div>
  </div>
</div>
  </div>
  </div>
</div>
  </div>
            </aside>
                      </div>

        </div>
        <!-- End main content -->

      </div>
      <!-- End content wrapper -->

      <!-- Start footer -->
      <footer id="nbcsports-footer">
          <div class="region region-footer">
    <div id="block-nbcs-footer-nbcs-footer" class="block block-nbcs-footer">


  <div class="content">
    <div id="footer-left">
  <div id="footer-logo"></div>
  <div id="footer-social-links">
    <span id="follow-us">FOLLOW US</span>
          <a href="http://facebook.com/NBCSports" class="genericon genericon-facebook">Facebook</a>          <a href="http://twitter.com/nbcsports" class="genericon genericon-twitter">Twitter</a>          <a href="http://plus.google.com/+NBCSports" class="genericon genericon-googleplus">Google Plus</a>          <a href="http://instagram.com/NBCsports" class="genericon genericon-instagram">Instagram</a>          <a href="/node/266336" id="directory">NBC Sports Social Directory</a>      </div>
  <div id="copyright">©2016 NBC Universal</div>
  <div id="footer-site-links">
          <a href="http://info.evidon.com/pub_info/267?v=1">Ad Choices</a>          <a href="https://together.nbcuni.com/">Advertise</a>          <a href="http://corporate.comcast.com/news-information/nbcuniversal-transaction/independent-programming">Independent Programming Report</a>          <a href="http://www.nbcuni.com/privacy/">Privacy Policy</a>          <a href="http://www.workinsports.com/nbc-sports-jobs/?ref=1525">Sports Jobs</a>          <a href="http://bit.ly/1UPP5G7">Terms of Use</a>      </div>
</div>
<div id="footer-right">
  <div id="footer-newsletter-link">
      <span id="newsletter-banner">GET NEWSLETTERS &amp; ALERTS</span>    <a href="http://www.nbcsports.com/newsletters/?cid=NL-FTR" id="newsletter-link" target="_blank">SIGN UP</a>   </div>
</div>
  </div>
</div>
<div id="block-bean-css-override-rracela" class="block block-bean block-bean-custom-block block-type-custom-block block-type-custom-block-single">


  <div class="content">
    <div class="entity entity-bean bean-custom-block clearfix">

  <div class="content">
    <div class="field field-name-field-content field-type-text-long field-label-hidden"><div class="field-items"><div class="field-item even"><style>
.breaking-news-marquee-message :only-child {
color: #FFFFFF;
}

#nbcsports-secondary-nav a.gold-nav {
font-size: 0;
letter-spacing: initial;
display: inline-block;
vertical-align: top;
padding: 0;
margin: 13px 12px;
height: 15px;
background: url(http://digitalassets.nbcsports.com/gold/nbcs-gold-gray.png) no-repeat 0 0;
width: 105px;

}

</style></div></div></div>  </div>
</div>
  </div>
</div>
  </div>
      </footer>
      <!-- End footer -->

    </div><!-- end #outer -->
  </div><!-- end #wrapper -->
  <!-- SiteCatalyst code version: 1.3
Copyright 1996-2016 Adobe, Inc. -->
<script type="text/javascript" src="/profiles/nbcs/modules/custom/nbcs_sitecatalyst/js/nbcs_adobe.js?oedmm1"></script>
<script type="text/javascript"><!--

/************* DO NOT ALTER ANYTHING BELOW THIS LINE ! **************/
var s_code=s.t();if(s_code)document.write(s_code)//--></script>
<script type="text/javascript"><!--
if(navigator.appVersion.indexOf('MSIE')>=0)document.write(unescape('%3C')+'\!-'+'-')
//--></script>
<noscript><img src="http://examplecom.112.2O7.net/b/ss/examplecom/1/H.20.3--NS/0/8565943" height="1" width="1" alt=""></noscript>
<!--/DO NOT REMOVE/-->
<!-- End SiteCatalyst code version: 1.3 -->
<script type="text/javascript">
<!--//--><![CDATA[//><!--
typeof(mps.writeFooter)=='function' && mps.writeFooter();
//--><!]]>
</script>
<script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"beacon":"bam.nr-data.net","licenseKey":"88172c94d3","applicationID":"15720345","transactionName":"YVAHZxEDCxYDW0BdVlgaJFAXCwoLTVZWV0ppWQxFBj0AHRZKVWtJV1IAbAsXBw==","queueTime":0,"applicationTime":507,"atts":"TRcEEVkZGBg=","errorBeacon":"bam.nr-data.net","agent":""}</script></body>
</html>
`
