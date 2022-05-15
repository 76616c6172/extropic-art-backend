(function(){"use strict";var t={3638:function(t,e,s){var n=s(9242),o=s(7154),r=s(7139),a=s(6265),i=s.n(a);const c="https://exia.art/api/0",l={jobs:[],selectedJob:[]},u={getJobs:t=>t.jobs,getSelectedJob:t=>t.selectedJob},d={async fetchJobs({commit:t}){try{return await i().get(`${c}/status`).then((e=>{if(200==e.status){const s=e.data.completed_jobs;t("FETCH_JOBS",s)}}))}catch(e){console.log(e)}},async sendNewJob({commit:t},e){try{return await i().post(`${c}/jobs`,e).then((e=>{if(200==e.status){const s=e.data;t("SEND_NEW_JOB",s)}}))}catch(s){console.log(s)}},getSelectedJob({commit:t},e){const s=this.getters.getJobs.filter((t=>t.jobid===e))[0];s.img=`${c}/img?${e}`,t("FETCH_SELECTED_JOB",s)},async getSelectedImg(){let t=l.selectedJob.img_path;try{return await i().get(`${t}`,{responseType:"blob"}).then((t=>{if(200==t.status)return new Promise((e=>{const s=t.data;e(s)}))}))}catch(e){console.log(e)}}},m={FETCH_JOBS(t,e){t.jobs=e},FETCH_SELECTED_JOB(t,e){t.selectedJob=e},SEND_NEW_JOB(t,e){t.jobs.push(e)}},g={state:l,mutations:m,actions:d,getters:u};var b=g,p=(0,r.MT)({modules:{api:b}}),v=s(678),h=s(3396);const f={class:"container-fluid"},_={class:"mt-5"},w={class:"row justify-content-center"},y={class:"col-lg-6 col-sm-12"},j={class:"row justify-content-center"},x={class:"col-lg-6 col-sm-12"},k={class:"row justify-content-center bg-light"},J={class:"col-lg-6 col-sm-12"},S={class:"row justify-content-center bg-light"},C={class:"col-lg-6 col-sm-12"},O={class:"row justify-content-center"},P={class:"col-lg-6 col-sm-12"};function E(t,e,s,n,o,r){const a=(0,h.up)("Typing"),i=(0,h.up)("Instructions"),c=(0,h.up)("ItemList"),l=(0,h.up)("Prompt"),u=(0,h.up)("Image");return(0,h.wg)(),(0,h.iD)("div",f,[(0,h._)("div",_,[(0,h._)("div",w,[(0,h._)("div",y,[(0,h.Wm)(a)])]),(0,h._)("div",j,[(0,h._)("div",x,[(0,h.Wm)(i)])]),(0,h._)("div",k,[(0,h._)("div",J,[(0,h.Wm)(c)])]),(0,h._)("div",S,[(0,h._)("div",C,[(0,h.Wm)(l)])]),(0,h._)("div",O,[(0,h._)("div",P,[(0,h.Wm)(u)])])])])}var T=s(2268);const D=(0,h._)("h1",{class:"text-start mb-4"},[(0,h._)("strong",null," High-resolution images generated by Ai ")],-1),I={class:"text-start"},L=(0,h._)("span",{class:"blink"},"|",-1);function N(t,e,s,n,o,r){return(0,h.wg)(),(0,h.iD)("div",null,[D,(0,h._)("p",I,[(0,h.Uk)((0,T.zw)(r.getText())+" ",1),L])])}var U={name:"TypingComponent",data(){return{i:0,txtOut:"",txtIn:"We leverage an AI Image generating technique called CLIP-Guided Diffusion to allow you to create compelling and beautiful images from just text inputs. Made with love by Zen and Valar!."}},methods:{getText(){return this.txtOut},delay(t){return new Promise((e=>setTimeout(e,t)))},async setText(){this.i<=this.txtIn.length&&(this.txtOut+=this.txtIn.charAt(this.i),await this.delay(40),this.i++,this.setText())}},mounted(){this.setText()}},W=s(89);const $=(0,W.Z)(U,[["render",N]]);var H=$;const F={class:"mb-5"},Z={class:"text-start"},q={key:0,class:"mt-5 mb-5"},B=(0,h._)("h2",{class:"text-start"},[(0,h._)("strong",null," How to ")],-1),z={class:"row"},R={class:"col-sm-1 col-lg-1 mt-2 mb-3"},M={class:"text-start"},Y={class:"col-sm-11 col-lg-5 mt-2 mb-3"},A={class:"text-start"},Q={key:0,class:"d-lg-block d-sm-none d-xs-none"};function V(t,e,s,n,o,r){return(0,h.wg)(),(0,h.iD)("div",F,[(0,h._)("div",Z,[(0,h._)("button",{onClick:e[0]||(e[0]=t=>this.visible=!this.visible),type:"button",class:"btn btn-outline-secondary"}," How to ")]),o.visible?((0,h.wg)(),(0,h.iD)("div",q,[B,(0,h._)("div",z,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(o.instructions,((t,e)=>((0,h.wg)(),(0,h.iD)(h.HY,{key:e},[(0,h._)("div",R,[(0,h._)("h4",M,"0"+(0,T.zw)(e+1),1)]),(0,h._)("div",Y,[(0,h._)("h5",A,(0,T.zw)(t.content),1)]),1==e||3==e?((0,h.wg)(),(0,h.iD)("hr",Q)):(0,h.kq)("",!0)],64)))),128))])])):(0,h.kq)("",!0)])}var K={name:"InstructionsComponent",data(){return{instructions:[{content:"Enter search term"},{content:"Click generate or hit enter"},{content:"Wait the image to be finished"},{content:"Enjoy and feel energized"}],visible:!1}}};const G=(0,W.Z)(K,[["render",V]]);var X=G;const tt={class:"pb-5"},et={class:"input-group"};function st(t,e,s,o,r,a){return(0,h.wg)(),(0,h.iD)("div",tt,[(0,h._)("div",et,[(0,h.wy)((0,h._)("input",{onKeyup:e[0]||(e[0]=(0,n.D2)((t=>a.onClickSendNewJob()),["enter"])),"onUpdate:modelValue":e[1]||(e[1]=t=>r.vPrompt=t),type:"text",class:"form-control",placeholder:"Enter your prompt","aria-label":"Enter your prompt"},null,544),[[n.nr,r.vPrompt]]),(0,h._)("button",{onClick:e[2]||(e[2]=t=>a.onClickSendNewJob()),class:"btn btn-outline-secondary",type:"submit"}," Generate ")])])}var nt={name:"PromptComponent",data(){return{vPrompt:""}},methods:{onClickSendNewJob(){var t=/^\s+$/;""==this.vPrompt||this.vPrompt.match(t)||(this.$store.dispatch("sendNewJob",{prompt:this.vPrompt}),this.vPrompt="")}}};const ot=(0,W.Z)(nt,[["render",st]]);var rt=ot;const at={class:"mt-5 mb-4"},it={class:"row"},ct=(0,h._)("h2",{class:"text-start mt-5 mb-4"},[(0,h._)("strong",null," ImageList")],-1),lt={class:"col-4 mb-3 mb-lg-0"},ut={class:"list-group-flush",style:{"padding-left":"0 !important"}};function dt(t,e,s,o,r,a){const i=(0,h.up)("Item");return(0,h.wg)(),(0,h.iD)("div",at,[(0,h._)("div",it,[ct,(0,h._)("form",lt,[(0,h.wy)((0,h._)("input",{"onUpdate:modelValue":e[0]||(e[0]=t=>r.searchQuery=t),type:"search",class:"form-control",placeholder:"Search...","aria-label":"Search"},null,512),[[n.nr,r.searchQuery]])])]),(0,h._)("ul",ut,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(a.getFilteredJobs,((t,e)=>((0,h.wg)(),(0,h.j4)(i,{key:e,job:t},null,8,["job"])))),128))])])}const mt=["disabled"],gt={class:"row"},bt={class:"col-lg-10 col-md-10"},pt={class:"col-lg-2 col-md-2"},vt={class:"progress mt-1"},ht=["aria-valuenow"];function ft(t,e,s,n,o,r){return(0,h.wg)(),(0,h.iD)("li",{disabled:"completed"!=s.job.job_status,class:(0,T.C_)([["completed"!=s.job.job_status?"disabled":""],"list-group-item list-group-item-action"])},[(0,h._)("div",gt,[(0,h._)("div",bt,[(0,h._)("p",{onClick:e[0]||(e[0]=t=>r.onClickSetSelected(t)),class:"text-start"},(0,T.zw)(s.job.prompt),1)]),(0,h._)("div",pt,[(0,h._)("div",{class:(0,T.C_)([r.getJobBorderClass,"badge border text-secondary"])},(0,T.zw)(s.job.job_status),3),(0,h._)("div",vt,[(0,h._)("div",{style:(0,T.j5)(`width: ${r.getProgressbarPercent}%;`),class:"progress-bar progress-bar-animated",role:"progressbar","aria-valuenow":r.getProgressbarPercent,"aria-valuemin":"0","aria-valuemax":"100"},(0,T.zw)(`${r.getProgressbarPercent}%`),13,ht)])])])],10,mt)}var _t={name:"ItemComponent",props:{job:{type:Object,required:!0,iteration_max:{type:String,required:!0},iteration_status:{type:String,required:!0},job_status:{type:String,required:!0},job_id:{type:String,required:!0},prompt:{type:String,required:!0}}},methods:{onClickSetSelected(t){let e=t.originalTarget.parentElement.parentElement.parentElement,s=document.getElementsByClassName("list-group-flush")[0].children,n="item-group-active";for(let o=0;o<s.length;o++)s[o].classList.remove(n);e.classList.add(n),"completed"!=this.job.job_status&&t.preventDefault(),this.$store.dispatch("getSelectedJob",this.job.jobid)}},computed:{getSelectedJob(){return this.$store.getters.getSelectedJob},getJobBorderClass(){let t;switch(this.job.job_status){case"completed":t="border-success";break;case"processing":t="border-info";break;case"queued":t="border-warning";break;default:break}return t},getProgressbarPercent(){return this.job.iteration_status/this.job.iteration_max*100}}};const wt=(0,W.Z)(_t,[["render",ft],["__scopeId","data-v-4cde9de2"]]);var yt=wt,jt={name:"StatusListComponent",components:{Item:yt},data(){return{searchQuery:""}},methods:{getFoundJobs(t){return t.filter((t=>-1!=t.prompt.toLowerCase().indexOf(this.searchQuery.toLowerCase())))}},computed:{getJobs(){return this.$store.getters.getJobs},getFilteredJobs(){let t=this.getJobs;return t.sort((t=>"completed"==t.job_status)),this.getFoundJobs(t)}},watch:{getFilteredJobs:{handler(t){t&&t.forEach((t=>{"accepted"==t.job_status&&setTimeout((()=>{this.$store.dispatch("fetchJobs")}),1500)}))}}},async mounted(){this.$store.dispatch("fetchJobs")}};const xt=(0,W.Z)(jt,[["render",dt]]);var kt=xt;const Jt={class:"mb-5 mt-5"},St={key:0,class:"spinner-border text-secondary",role:"status"},Ct=(0,h._)("span",{class:"visually-hidden"},"Loading...",-1),Ot=[Ct],Pt=["src"];function Et(t,e,s,n,o,r){return(0,h.wg)(),(0,h.iD)("div",null,[(0,h._)("div",Jt,[o.isLoading?((0,h.wg)(),(0,h.iD)("div",St,Ot)):(0,h.kq)("",!0),(0,h._)("img",{src:o.imgObjectURL,class:"img-fluid img-thumbnail",alt:""},null,8,Pt)])])}var Tt={name:"ImageComponent",data(){return{imgObjectURL:"https://via.placeholder.com/1920x1024.png?text=This%20is%20zen%27s%20placeholder",isLoading:!1}},methods:{createImgObjectURL(){this.isLoading=!0,this.$store.dispatch("getSelectedImg").then((t=>{var e=new Image;let s=new Blob([t],{type:"image/png"}),n=URL.createObjectURL(s);e.src=n,this.imgObjectURL=e.src})).finally((()=>{this.isLoading=!1}))}},computed:{getSelectedJob(){return this.$store.getters.getSelectedJob}},watch:{getSelectedJob:{handler(){this.createImgObjectURL()}}}};const Dt=(0,W.Z)(Tt,[["render",Et]]);var It=Dt,Lt={name:"HomeComponent",components:{Typing:H,Instructions:X,Prompt:rt,ItemList:kt,Image:It},props:{msg:String}};const Nt=(0,W.Z)(Lt,[["render",E]]);var Ut=Nt;const Wt=[{path:"/",name:"Home",component:Ut}],$t=(0,v.p7)({history:(0,v.PO)("/"),routes:Wt});var Ht=$t;const Ft={class:"appContainer"},Zt={class:"row"},qt={class:"col-12"};function Bt(t,e,s,n,o,r){const a=(0,h.up)("Navbar"),i=(0,h.up)("Home"),c=(0,h.up)("Footer");return(0,h.wg)(),(0,h.iD)("div",Ft,[(0,h.Wm)(a),(0,h._)("div",Zt,[(0,h._)("div",qt,[(0,h.Wm)(i)])]),(0,h.Wm)(c)])}const zt={class:"navbar navbar-expand-lg navbar-light fixed-top"},Rt={class:"container"},Mt=(0,h._)("a",{class:"navbar-brand",href:"#"},"Exia",-1),Yt=(0,h._)("button",{class:"navbar-toggler",type:"button","data-bs-toggle":"collapse","data-bs-target":"#navbarNav","aria-controls":"navbarNav","aria-expanded":"false","aria-label":"Toggle navigation",style:{width:"auto"}},[(0,h._)("span",{class:"navbar-toggler-icon"})],-1),At={class:"collapse navbar-collapse",id:"navbarNav"},Qt={class:"navbar-nav"},Vt={class:"nav-item"},Kt=(0,h.Uk)("Home"),Gt={class:"nav-item"},Xt=(0,h.Uk)("Settings"),te={class:"nav-item"},ee=(0,h.Uk)("About"),se={class:"nav-item"},ne=(0,h.Uk)("Login");function oe(t,e,s,n,o,r){const a=(0,h.up)("router-link");return(0,h.wg)(),(0,h.iD)("nav",zt,[(0,h._)("div",Rt,[Mt,Yt,(0,h._)("div",At,[(0,h._)("ul",Qt,[(0,h._)("li",Vt,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start active"},{default:(0,h.w5)((()=>[Kt])),_:1})]),(0,h._)("div",Gt,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[Xt])),_:1})]),(0,h._)("div",te,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[ee])),_:1})]),(0,h._)("div",se,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[ne])),_:1})])])])])])}var re={name:"NavbarComponent"};const ae=(0,W.Z)(re,[["render",oe]]);var ie=ae;const ce={class:"footer text-muted mt-auto py-3 bg-light"},le={class:"text-center p-3"},ue=(0,h.Uk)(" Made with "),de=(0,h._)("i",{class:"fa fa-heart","aria-hidden":"true"},null,-1);function me(t,e,s,n,o,r){return(0,h.wg)(),(0,h.iD)("footer",ce,[(0,h._)("div",le,[ue,de,(0,h.Uk)(" by Valar and Zendo in "+(0,T.zw)(r.getYear),1)])])}var ge={name:"FooterComponent",data(){return{}},computed:{getYear(){return(new Date).getFullYear()}}};const be=(0,W.Z)(ge,[["render",me]]);var pe=be,ve={name:"App",components:{Home:Ut,Navbar:ie,Footer:pe}};const he=(0,W.Z)(ve,[["render",Bt]]);var fe=he;let _e=(0,n.ri)(fe);_e.use(o),_e.use(p),_e.use(Ht),_e.mount("#app")}},e={};function s(n){var o=e[n];if(void 0!==o)return o.exports;var r=e[n]={exports:{}};return t[n].call(r.exports,r,r.exports,s),r.exports}s.m=t,function(){var t=[];s.O=function(e,n,o,r){if(!n){var a=1/0;for(u=0;u<t.length;u++){n=t[u][0],o=t[u][1],r=t[u][2];for(var i=!0,c=0;c<n.length;c++)(!1&r||a>=r)&&Object.keys(s.O).every((function(t){return s.O[t](n[c])}))?n.splice(c--,1):(i=!1,r<a&&(a=r));if(i){t.splice(u--,1);var l=o();void 0!==l&&(e=l)}}return e}r=r||0;for(var u=t.length;u>0&&t[u-1][2]>r;u--)t[u]=t[u-1];t[u]=[n,o,r]}}(),function(){s.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return s.d(e,{a:e}),e}}(),function(){s.d=function(t,e){for(var n in e)s.o(e,n)&&!s.o(t,n)&&Object.defineProperty(t,n,{enumerable:!0,get:e[n]})}}(),function(){s.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){s.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){s.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})}}(),function(){var t={143:0};s.O.j=function(e){return 0===t[e]};var e=function(e,n){var o,r,a=n[0],i=n[1],c=n[2],l=0;if(a.some((function(e){return 0!==t[e]}))){for(o in i)s.o(i,o)&&(s.m[o]=i[o]);if(c)var u=c(s)}for(e&&e(n);l<a.length;l++)r=a[l],s.o(t,r)&&t[r]&&t[r][0](),t[r]=0;return s.O(u)},n=self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[];n.forEach(e.bind(null,0)),n.push=e.bind(null,n.push.bind(n))}();var n=s.O(void 0,[998],(function(){return s(3638)}));n=s.O(n)})();
//# sourceMappingURL=app.eb69d824.js.map