(function(){"use strict";var t={9314:function(t,e,o){var s=o(9242),n=o(7154),a=o(7139),r=o(6265),i=o.n(r);const l="https://exia.art/api/0",c={jobs:[],selectedJob:[]},u={getJobs:t=>t.jobs,getSelectedJob:t=>t.selectedJob},d={async fetchJobs({commit:t}){try{return await i().get(`${l}/status`).then((e=>{if(200==e.status){const o=e.data.completed_jobs;t("FETCH_JOBS",o)}}))}catch(e){console.log(e)}},async sendNewJob({commit:t},e){try{return await i().post(`${l}/jobs`,e).then((e=>{if(200==e.status){const o=e.data;t("SEND_NEW_JOB",o)}}))}catch(o){console.log(o)}},getSelectedJob({commit:t},e){const o=this.getters.getJobs.filter((t=>t.jobid===e))[0];o.img=`${l}/img?${e}`,t("FETCH_SELECTED_JOB",o)}},p={FETCH_JOBS(t,e){t.jobs=e},FETCH_SELECTED_JOB(t,e){t.selectedJob=e},SEND_NEW_JOB(t,e){t.jobs.push(e),console.log("SEND_NEW_NOB"),console.log(t.jobs)}},m={state:c,mutations:p,actions:d,getters:u};var b=m,g=(0,a.MT)({modules:{api:b}}),v=o(678),h=o(3396);const f={class:"container-fluid"},_={class:"mt-5"},w={class:"row justify-content-center"},y={class:"col-lg-6 col-sm-12"},x={class:"row justify-content-center"},k={class:"col-lg-6 col-sm-12"},S={class:"row justify-content-center bg-light"},j={class:"col-lg-6 col-sm-12"},C={class:"row justify-content-center bg-light"},J={class:"col-lg-6 col-sm-12"},D={class:"row justify-content-center bg-light"},E={class:"col-lg-6 col-sm-12"};function O(t,e,o,s,n,a){const r=(0,h.up)("Typing"),i=(0,h.up)("Instructions"),l=(0,h.up)("StatusItemList"),c=(0,h.up)("Prompt"),u=(0,h.up)("Carousel");return(0,h.wg)(),(0,h.iD)("div",f,[(0,h._)("div",_,[(0,h._)("div",w,[(0,h._)("div",y,[(0,h.Wm)(r)])]),(0,h._)("div",x,[(0,h._)("div",k,[(0,h.Wm)(i)])]),(0,h._)("div",S,[(0,h._)("div",j,[(0,h.Wm)(l)])]),(0,h._)("div",C,[(0,h._)("div",J,[(0,h.Wm)(c)])]),(0,h._)("div",D,[(0,h._)("div",E,[(0,h.Wm)(u)])])])])}var P=o(2268);const T=(0,h._)("h1",{class:"text-start mb-4"},[(0,h._)("strong",null," High-resolution images generated by Ai ")],-1),I={class:"text-start"},H=(0,h._)("span",{class:"blink"},"|",-1);function B(t,e,o,s,n,a){return(0,h.wg)(),(0,h.iD)("div",null,[T,(0,h._)("p",I,[(0,h.Uk)((0,P.zw)(a.getText())+" ",1),H])])}var N={name:"TypingComponent",data(){return{i:0,txtOut:"",txtIn:"We leverage an AI Image generating technique called CLIP-Guided Diffusion to allow you to create compelling and beautiful images from just text inputs. Made with love by Zen and Valar!."}},methods:{getText(){return this.txtOut},delay(t){return new Promise((e=>setTimeout(e,t)))},async setText(){this.i<=this.txtIn.length&&(this.txtOut+=this.txtIn.charAt(this.i),await this.delay(40),this.i++,this.setText())}},mounted(){this.setText()}},W=o(89);const L=(0,W.Z)(N,[["render",B]]);var $=L;const M={class:"mb-5"},Z={class:"text-start"},q={key:0,class:"mt-5 mb-5"},z=(0,h._)("h2",{class:"text-start"},[(0,h._)("strong",null," How to ")],-1),F={class:"row"},U={class:"col-sm-1 col-lg-1 mt-2 mb-3"},V={class:"text-start"},Y={class:"col-sm-11 col-lg-5 mt-2 mb-3"},A={class:"text-start"},Q={key:0,class:"d-lg-block d-sm-none d-xs-none"};function K(t,e,o,s,n,a){return(0,h.wg)(),(0,h.iD)("div",M,[(0,h._)("div",Z,[(0,h._)("button",{onClick:e[0]||(e[0]=t=>this.visible=!this.visible),type:"button",class:"btn btn-outline-secondary"}," How to ")]),n.visible?((0,h.wg)(),(0,h.iD)("div",q,[z,(0,h._)("div",F,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(n.instructions,((t,e)=>((0,h.wg)(),(0,h.iD)(h.HY,{key:e},[(0,h._)("div",U,[(0,h._)("h4",V,"0"+(0,P.zw)(e+1),1)]),(0,h._)("div",Y,[(0,h._)("h5",A,(0,P.zw)(t.content),1)]),1==e||3==e?((0,h.wg)(),(0,h.iD)("hr",Q)):(0,h.kq)("",!0)],64)))),128))])])):(0,h.kq)("",!0)])}var G={name:"InstructionsComponent",data(){return{instructions:[{content:"Enter search term"},{content:"Click generate or hit enter"},{content:"Wait the image to be finished"},{content:"Enjoy and feel energized"}],visible:!1}}};const R=(0,W.Z)(G,[["render",K]]);var X=R;const tt={class:"input-group"};function et(t,e,o,n,a,r){return(0,h.wg)(),(0,h.iD)("div",null,[(0,h._)("div",tt,[(0,h.wy)((0,h._)("input",{onKeyup:e[0]||(e[0]=(0,s.D2)((t=>r.onClickSendNewJob()),["enter"])),"onUpdate:modelValue":e[1]||(e[1]=t=>a.vPrompt=t),type:"text",class:"form-control",placeholder:"Enter your prompt","aria-label":"Enter your prompt"},null,544),[[s.nr,a.vPrompt]]),(0,h._)("button",{onClick:e[2]||(e[2]=t=>r.onClickSendNewJob()),class:"btn btn-outline-secondary",type:"submit"}," Generate ")])])}var ot={name:"PromptComponent",data(){return{vPrompt:""}},methods:{onClickSendNewJob(){this.$store.dispatch("sendNewJob",{prompt:this.vPrompt})}}};const st=(0,W.Z)(ot,[["render",et]]);var nt=st;const at={class:"mt-5 mb-4"},rt={class:"row"},it=(0,h._)("h2",{class:"text-start mt-5 mb-4"},[(0,h._)("strong",null," ImageList")],-1),lt={class:"col-4 mb-3 mb-lg-0"},ct={class:"list-group-flush",style:{"padding-left":"0 !important"}};function ut(t,e,o,n,a,r){const i=(0,h.up)("StatusItem");return(0,h.wg)(),(0,h.iD)("div",at,[(0,h._)("div",rt,[it,(0,h._)("form",lt,[(0,h.wy)((0,h._)("input",{"onUpdate:modelValue":e[0]||(e[0]=t=>a.searchQuery=t),type:"search",class:"form-control",placeholder:"Search...","aria-label":"Search"},null,512),[[s.nr,a.searchQuery]])])]),(0,h._)("ul",ct,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(r.getSearchedProducts,((t,e)=>((0,h.wg)(),(0,h.j4)(i,{key:e,job:t},null,8,["job"])))),128))])])}const dt={class:"row"},pt={class:"col-lg-10 col-md-10"},mt={class:"text-start"},bt={class:"col-lg-2 col-md-2"},gt={class:"badge border text-secondary"};function vt(t,e,o,s,n,a){const r=(0,h.up)("ConfirmDialogue");return(0,h.wg)(),(0,h.iD)(h.HY,null,[(0,h._)("li",{onClick:e[0]||(e[0]=t=>a.onClickSetSelected(o.job.jobid)),class:"list-group-item list-group-item-action"},[(0,h._)("div",dt,[(0,h._)("div",pt,[(0,h._)("p",mt,(0,P.zw)(o.job.prompt),1)]),(0,h._)("div",bt,[(0,h._)("span",{class:(0,P.C_)([a.getJobBorderClass(o.job.job_status),"badge border text-secondary"])},(0,P.zw)(o.job.job_status),3),(0,h._)("span",gt,(0,P.zw)(o.job.iteration_status)+"/"+(0,P.zw)(o.job.iteration_max),1)])])]),(0,h.Wm)(r,{ref:"confirmDialogue"},null,512)],64)}const ht=["innerHTML"],ft=["innerHTML"],_t=["src"],wt={class:"popupmodal__buttons"};function yt(t,e,o,s,n,a){const r=(0,h.up)("popup-modal");return(0,h.wg)(),(0,h.j4)(r,{ref:"popup"},{default:(0,h.w5)((()=>[(0,h._)("h2",{class:"popupmodal__title",style:{"margin-top":"0"},innerHTML:t.title},null,8,ht),(0,h._)("p",{class:"popupmodal__message",innerHTML:t.message},null,8,ft),(0,h._)("img",{class:"img-fluid",src:`${t.image}`,alt:""},null,8,_t),(0,h._)("div",wt,[(0,h._)("button",{class:"popupmodal__buttons-cancel",onClick:e[0]||(e[0]=(...t)=>a._cancel&&a._cancel(...t))},(0,P.zw)(t.cancelButton),1),(0,h._)("button",{class:"popupmodal__buttons-ok",onClick:e[1]||(e[1]=(...t)=>a._confirm&&a._confirm(...t))},(0,P.zw)(t.okButton),1)])])),_:1},512)}const xt={key:0,class:"popupmodal"},kt={class:"popupmodal__window"};function St(t,e,o,n,a,r){return(0,h.wg)(),(0,h.j4)(s.uT,{name:"fade"},{default:(0,h.w5)((()=>[t.isVisible?((0,h.wg)(),(0,h.iD)("div",xt,[(0,h._)("div",kt,[(0,h.WI)(t.$slots,"default")])])):(0,h.kq)("",!0)])),_:3})}var jt={name:"PopupModal",data:()=>({isVisible:!1}),methods:{open(){this.isVisible=!0},close(){this.isVisible=!1}}};const Ct=(0,W.Z)(jt,[["render",St]]);var Jt=Ct,Dt={name:"ConfirmDialogue",components:{PopupModal:Jt},data:()=>({title:void 0,message:void 0,image:void 0,okButton:void 0,cancelButton:"Cancel",resolvePromise:void 0,rejectPromise:void 0}),methods:{show(t={}){return this.title=t.title,this.message=t.message,this.image=t.image,this.okButton=t.okButton,t.cancelButton&&(this.cancelButton=t.cancelButton),this.$refs.popup.open(),new Promise(((t,e)=>{this.resolvePromise=t,this.rejectPromise=e}))},_confirm(){this.$refs.popup.close(),this.resolvePromise(!0)},_cancel(){this.$refs.popup.close(),this.resolvePromise(!1)}}};const Et=(0,W.Z)(Dt,[["render",yt]]);var Ot=Et,Pt={name:"StatusItemComponent",components:{ConfirmDialogue:Ot},props:{job:{type:Object,required:!0,iteration_max:{type:String,required:!0},iteration_status:{type:String,required:!0},job_status:{type:String,required:!0},job_id:{type:String,required:!0},prompt:{type:String,required:!0}}},methods:{onClickSetSelected(t){this.$store.dispatch("getSelectedJob",t),this.buildModalDialogue()},async buildModalDialogue(){console.log("triggered");await this.$refs.confirmDialogue.show({title:this.getSelectedJob.prompt,image:this.getSelectedJob.img,message:"",okButton:"Löschen"})},getJobBorderClass(t){let e;switch(t){case"completed":e="border-success";break;case"processing":e="border-info";break;case"queued":e="border-warning";break;default:break}return e}},computed:{getSelectedJob(){return this.$store.getters.getSelectedJob}}};const Tt=(0,W.Z)(Pt,[["render",vt],["__scopeId","data-v-269dde5a"]]);var It=Tt,Ht={name:"StatusListComponent",components:{StatusItem:It},data(){return{jobs:[],searchQuery:""}},methods:{loadJobs(t){0==this.jobs.length&&(t.forEach((t=>{""!=t.prompt&&this.jobs.push(t)})),this.jobs=this.jobs.sort((t=>"completed"==t.job_status)))}},computed:{getJobsFromStore(){return this.$store.getters.getJobs},getSearchedProducts(){return this.jobs.filter((t=>-1!=t.prompt.toLowerCase().indexOf(this.searchQuery.toLowerCase())))}},watch:{getJobsFromStore:{handler(t){t&&this.loadJobs(t)},immediate:!0}},async mounted(){this.$store.dispatch("fetchJobs")}};const Bt=(0,W.Z)(Ht,[["render",ut]]);var Nt=Bt;const Wt={id:"carouselExampleIndicators",class:"carousel slide","data-bs-ride":"carousel"},Lt=(0,h.uE)('<div class="carousel-indicators"><button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button><button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1" aria-label="Slide 2"></button><button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2" aria-label="Slide 3"></button></div><div class="carousel-inner"><div class="carousel-item active"><img src="https://exia.art/api/0/img?4" class="d-block w-100" alt="..."><div class="carousel-caption d-none d-md-block"><p>Some representative placeholder content for the first slide.</p></div></div><div class="carousel-item"><img src="https://exia.art/api/0/img?2" class="d-block w-100" alt="..."><div class="carousel-caption d-none d-md-block"><p>Some representative placeholder content for the first slide.</p></div></div><div class="carousel-item"><img src="https://exia.art/api/0/img?5" class="d-block w-100" alt="..."><div class="carousel-caption d-none d-md-block"><p>Some representative placeholder content for the first slide.</p></div></div></div><button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="prev"><span class="carousel-control-prev-icon" aria-hidden="true"></span><span class="visually-hidden">Previous</span></button><button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide="next"><span class="carousel-control-next-icon" aria-hidden="true"></span><span class="visually-hidden">Next</span></button>',4),$t=[Lt];function Mt(t,e,o,s,n,a){return(0,h.wg)(),(0,h.iD)("div",Wt,$t)}var Zt={name:"CarouselComponent"};const qt=(0,W.Z)(Zt,[["render",Mt]]);var zt=qt,Ft={name:"HomeComponent",components:{Typing:$,Instructions:X,Prompt:nt,StatusItemList:Nt,Carousel:zt},props:{msg:String}};const Ut=(0,W.Z)(Ft,[["render",O]]);var Vt=Ut;const Yt=[{path:"/",name:"Home",component:Vt}],At=(0,v.p7)({history:(0,v.PO)("/"),routes:Yt});var Qt=At;const Kt={class:"appContainer"},Gt={class:"row"},Rt={class:"col-12"};function Xt(t,e,o,s,n,a){const r=(0,h.up)("Navbar"),i=(0,h.up)("Home"),l=(0,h.up)("Footer");return(0,h.wg)(),(0,h.iD)("div",Kt,[(0,h.Wm)(r),(0,h._)("div",Gt,[(0,h._)("div",Rt,[(0,h.Wm)(i)])]),(0,h.Wm)(l)])}const te={class:"navbar navbar-expand-lg navbar-light fixed-top"},ee={class:"container"},oe=(0,h._)("a",{class:"navbar-brand",href:"#"},"Exia",-1),se=(0,h._)("button",{class:"navbar-toggler",type:"button","data-bs-toggle":"collapse","data-bs-target":"#navbarNav","aria-controls":"navbarNav","aria-expanded":"false","aria-label":"Toggle navigation",style:{width:"auto"}},[(0,h._)("span",{class:"navbar-toggler-icon"})],-1),ne={class:"collapse navbar-collapse",id:"navbarNav"},ae={class:"navbar-nav"},re={class:"nav-item"},ie=(0,h.Uk)("Home"),le={class:"nav-item"},ce=(0,h.Uk)("Settings"),ue={class:"nav-item"},de=(0,h.Uk)("About"),pe={class:"nav-item"},me=(0,h.Uk)("Login");function be(t,e,o,s,n,a){const r=(0,h.up)("router-link");return(0,h.wg)(),(0,h.iD)("nav",te,[(0,h._)("div",ee,[oe,se,(0,h._)("div",ne,[(0,h._)("ul",ae,[(0,h._)("li",re,[(0,h.Wm)(r,{to:"/",class:"nav-link text-start active"},{default:(0,h.w5)((()=>[ie])),_:1})]),(0,h._)("div",le,[(0,h.Wm)(r,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[ce])),_:1})]),(0,h._)("div",ue,[(0,h.Wm)(r,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[de])),_:1})]),(0,h._)("div",pe,[(0,h.Wm)(r,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[me])),_:1})])])])])])}var ge={name:"NavbarComponent"};const ve=(0,W.Z)(ge,[["render",be]]);var he=ve;const fe={class:"footer text-muted mt-auto py-3 bg-light"},_e={class:"text-center p-3"},we=(0,h.Uk)(" Made with "),ye=(0,h._)("i",{class:"fa fa-heart","aria-hidden":"true"},null,-1);function xe(t,e,o,s,n,a){return(0,h.wg)(),(0,h.iD)("footer",fe,[(0,h._)("div",_e,[we,ye,(0,h.Uk)(" by Valar and Zendo in "+(0,P.zw)(a.getYear),1)])])}var ke={name:"FooterComponent",data(){return{}},computed:{getYear(){return(new Date).getFullYear()}}};const Se=(0,W.Z)(ke,[["render",xe]]);var je=Se,Ce={name:"App",components:{Home:Vt,Navbar:he,Footer:je}};const Je=(0,W.Z)(Ce,[["render",Xt]]);var De=Je;let Ee=(0,s.ri)(De);Ee.use(n),Ee.use(g),Ee.use(Qt),Ee.mount("#app")}},e={};function o(s){var n=e[s];if(void 0!==n)return n.exports;var a=e[s]={exports:{}};return t[s].call(a.exports,a,a.exports,o),a.exports}o.m=t,function(){var t=[];o.O=function(e,s,n,a){if(!s){var r=1/0;for(u=0;u<t.length;u++){s=t[u][0],n=t[u][1],a=t[u][2];for(var i=!0,l=0;l<s.length;l++)(!1&a||r>=a)&&Object.keys(o.O).every((function(t){return o.O[t](s[l])}))?s.splice(l--,1):(i=!1,a<r&&(r=a));if(i){t.splice(u--,1);var c=n();void 0!==c&&(e=c)}}return e}a=a||0;for(var u=t.length;u>0&&t[u-1][2]>a;u--)t[u]=t[u-1];t[u]=[s,n,a]}}(),function(){o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,{a:e}),e}}(),function(){o.d=function(t,e){for(var s in e)o.o(e,s)&&!o.o(t,s)&&Object.defineProperty(t,s,{enumerable:!0,get:e[s]})}}(),function(){o.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})}}(),function(){var t={143:0};o.O.j=function(e){return 0===t[e]};var e=function(e,s){var n,a,r=s[0],i=s[1],l=s[2],c=0;if(r.some((function(e){return 0!==t[e]}))){for(n in i)o.o(i,n)&&(o.m[n]=i[n]);if(l)var u=l(o)}for(e&&e(s);c<r.length;c++)a=r[c],o.o(t,a)&&t[a]&&t[a][0](),t[a]=0;return o.O(u)},s=self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[];s.forEach(e.bind(null,0)),s.push=e.bind(null,s.push.bind(s))}();var s=o.O(void 0,[998],(function(){return o(9314)}));s=o.O(s)})();
//# sourceMappingURL=app.1a256fa3.js.map