(function(){"use strict";var t={1255:function(t,e,o){var s=o(9242),n=o(7154),r=o(7139),a=o(6265),i=o.n(a);const l="https://exia.art/api/0",c={jobs:[],selectedJob:[]},u={getJobs:t=>t.jobs,getSelectedJob:t=>t.selectedJob},d={async fetchJobs({commit:t}){try{return await i().get(`${l}/status`).then((e=>{if(200==e.status){const o=e.data.completed_jobs;t("FETCH_JOBS",o)}}))}catch(e){console.log(e)}},async sendNewJob({commit:t},e){try{return await i().post(`${l}/jobs`,e).then((e=>{if(200==e.status){const o=e.data;t("SEND_NEW_JOB",o)}}))}catch(o){console.log(o)}},getSelectedJob({commit:t},e){const o=this.getters.getJobs.filter((t=>t.jobid===e))[0];o.img=`${l}/img?${e}`,t("FETCH_SELECTED_JOB",o)}},m={FETCH_JOBS(t,e){t.jobs=e},FETCH_SELECTED_JOB(t,e){t.selectedJob=e},SEND_NEW_JOB(t,e){t.jobs.push(e)}},p={state:c,mutations:m,actions:d,getters:u};var g=p,b=(0,r.MT)({modules:{api:g}}),v=o(678),h=o(3396);const f={class:"container-fluid"},_={class:"mt-5"},w={class:"row justify-content-center"},y={class:"col-lg-6 col-sm-12"},j={class:"row justify-content-center"},k={class:"col-lg-6 col-sm-12"},x={class:"row justify-content-center bg-light"},S={class:"col-lg-6 col-sm-12"},C={class:"row justify-content-center bg-light"},J={class:"col-lg-6 col-sm-12"},P=(0,h._)("div",{class:"row justify-content-center bg-light"},[(0,h._)("div",{class:"col-lg-6 col-sm-12"})],-1);function D(t,e,o,s,n,r){const a=(0,h.up)("Typing"),i=(0,h.up)("Instructions"),l=(0,h.up)("StatusItemList"),c=(0,h.up)("Prompt");return(0,h.wg)(),(0,h.iD)("div",f,[(0,h._)("div",_,[(0,h._)("div",w,[(0,h._)("div",y,[(0,h.Wm)(a)])]),(0,h._)("div",j,[(0,h._)("div",k,[(0,h.Wm)(i)])]),(0,h._)("div",x,[(0,h._)("div",S,[(0,h.Wm)(l)])]),(0,h._)("div",C,[(0,h._)("div",J,[(0,h.Wm)(c)])]),P])])}var T=o(2268);const O=(0,h._)("h1",{class:"text-start mb-4"},[(0,h._)("strong",null," High-resolution images generated by Ai ")],-1),E={class:"text-start"},H=(0,h._)("span",{class:"blink"},"|",-1);function $(t,e,o,s,n,r){return(0,h.wg)(),(0,h.iD)("div",null,[O,(0,h._)("p",E,[(0,h.Uk)((0,T.zw)(r.getText())+" ",1),H])])}var B={name:"TypingComponent",data(){return{i:0,txtOut:"",txtIn:"We leverage an AI Image generating technique called CLIP-Guided Diffusion to allow you to create compelling and beautiful images from just text inputs. Made with love by Zen and Valar!."}},methods:{getText(){return this.txtOut},delay(t){return new Promise((e=>setTimeout(e,t)))},async setText(){this.i<=this.txtIn.length&&(this.txtOut+=this.txtIn.charAt(this.i),await this.delay(40),this.i++,this.setText())}},mounted(){this.setText()}},I=o(89);const W=(0,I.Z)(B,[["render",$]]);var L=W;const N={class:"mb-5"},F={class:"text-start"},M={key:0,class:"mt-5 mb-5"},Z=(0,h._)("h2",{class:"text-start"},[(0,h._)("strong",null," How to ")],-1),q={class:"row"},z={class:"col-sm-1 col-lg-1 mt-2 mb-3"},U={class:"text-start"},V={class:"col-sm-11 col-lg-5 mt-2 mb-3"},Y={class:"text-start"},A={key:0,class:"d-lg-block d-sm-none d-xs-none"};function Q(t,e,o,s,n,r){return(0,h.wg)(),(0,h.iD)("div",N,[(0,h._)("div",F,[(0,h._)("button",{onClick:e[0]||(e[0]=t=>this.visible=!this.visible),type:"button",class:"btn btn-outline-secondary"}," How to ")]),n.visible?((0,h.wg)(),(0,h.iD)("div",M,[Z,(0,h._)("div",q,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(n.instructions,((t,e)=>((0,h.wg)(),(0,h.iD)(h.HY,{key:e},[(0,h._)("div",z,[(0,h._)("h4",U,"0"+(0,T.zw)(e+1),1)]),(0,h._)("div",V,[(0,h._)("h5",Y,(0,T.zw)(t.content),1)]),1==e||3==e?((0,h.wg)(),(0,h.iD)("hr",A)):(0,h.kq)("",!0)],64)))),128))])])):(0,h.kq)("",!0)])}var K={name:"InstructionsComponent",data(){return{instructions:[{content:"Enter search term"},{content:"Click generate or hit enter"},{content:"Wait the image to be finished"},{content:"Enjoy and feel energized"}],visible:!1}}};const G=(0,I.Z)(K,[["render",Q]]);var R=G;const X={class:"input-group"};function tt(t,e,o,n,r,a){return(0,h.wg)(),(0,h.iD)("div",null,[(0,h._)("div",X,[(0,h.wy)((0,h._)("input",{onKeyup:e[0]||(e[0]=(0,s.D2)((t=>a.onClickSendNewJob()),["enter"])),"onUpdate:modelValue":e[1]||(e[1]=t=>r.vPrompt=t),type:"text",class:"form-control",placeholder:"Enter your prompt","aria-label":"Enter your prompt"},null,544),[[s.nr,r.vPrompt]]),(0,h._)("button",{onClick:e[2]||(e[2]=t=>a.onClickSendNewJob()),class:"btn btn-outline-secondary",type:"submit"}," Generate ")])])}var et={name:"PromptComponent",data(){return{vPrompt:""}},methods:{onClickSendNewJob(){var t=/^\s+$/;""==this.vPrompt||this.vPrompt.match(t)||(this.$store.dispatch("sendNewJob",{prompt:this.vPrompt}),this.vPrompt="")}}};const ot=(0,I.Z)(et,[["render",tt]]);var st=ot;const nt={class:"mt-5 mb-4"},rt={class:"row"},at=(0,h._)("h2",{class:"text-start mt-5 mb-4"},[(0,h._)("strong",null," ImageList")],-1),it={class:"col-4 mb-3 mb-lg-0"},lt={class:"list-group-flush",style:{"padding-left":"0 !important"}};function ct(t,e,o,n,r,a){const i=(0,h.up)("StatusItem");return(0,h.wg)(),(0,h.iD)("div",nt,[(0,h._)("div",rt,[at,(0,h._)("form",it,[(0,h.wy)((0,h._)("input",{"onUpdate:modelValue":e[0]||(e[0]=t=>r.searchQuery=t),type:"search",class:"form-control",placeholder:"Search...","aria-label":"Search"},null,512),[[s.nr,r.searchQuery]])])]),(0,h._)("ul",lt,[((0,h.wg)(!0),(0,h.iD)(h.HY,null,(0,h.Ko)(a.getFilteredJobs,((t,e)=>((0,h.wg)(),(0,h.j4)(i,{key:e,job:t},null,8,["job"])))),128))])])}const ut={class:"row"},dt={class:"col-lg-10 col-md-10"},mt={class:"text-start"},pt={class:"col-lg-2 col-md-2"},gt={class:"progress mt-1"},bt=["aria-valuenow"];function vt(t,e,o,s,n,r){const a=(0,h.up)("ConfirmDialogue");return(0,h.wg)(),(0,h.iD)(h.HY,null,[(0,h._)("li",{onClick:e[0]||(e[0]=t=>r.onClickSetSelected(o.job.jobid)),class:"list-group-item list-group-item-action"},[(0,h._)("div",ut,[(0,h._)("div",dt,[(0,h._)("p",mt,(0,T.zw)(o.job.prompt),1)]),(0,h._)("div",pt,[(0,h._)("div",{class:(0,T.C_)([r.getJobBorderClass(o.job.job_status),"badge border text-secondary"])},(0,T.zw)(o.job.job_status),3),(0,h._)("div",gt,[(0,h._)("div",{style:(0,T.j5)(`width: ${r.getProgressbarPercent(o.job.iteration_status,o.job.iteration_max)}%;`),class:"progress-bar progress-bar-animated",role:"progressbar","aria-valuenow":r.getProgressbarPercent(o.job.iteration_status,o.job.iteration_max),"aria-valuemin":"0","aria-valuemax":"100"},(0,T.zw)(`${r.getProgressbarPercent(o.job.iteration_status,o.job.iteration_max)}%`),13,bt)])])])]),(0,h.Wm)(a,{ref:"confirmDialogue"},null,512)],64)}const ht=["innerHTML"],ft=["innerHTML"],_t=["src"],wt={class:"popupmodal__buttons"};function yt(t,e,o,s,n,r){const a=(0,h.up)("popup-modal");return(0,h.wg)(),(0,h.j4)(a,{ref:"popup"},{default:(0,h.w5)((()=>[(0,h._)("h2",{class:"popupmodal__title",style:{"margin-top":"0"},innerHTML:t.title},null,8,ht),(0,h._)("p",{class:"popupmodal__message",innerHTML:t.message},null,8,ft),(0,h._)("img",{class:"img-fluid",src:`${t.image}`,alt:""},null,8,_t),(0,h._)("div",wt,[(0,h._)("button",{class:"popupmodal__buttons-cancel",onClick:e[0]||(e[0]=(...t)=>r._cancel&&r._cancel(...t))},(0,T.zw)(t.cancelButton),1),(0,h._)("button",{class:"popupmodal__buttons-ok",onClick:e[1]||(e[1]=(...t)=>r._confirm&&r._confirm(...t))},(0,T.zw)(t.okButton),1)])])),_:1},512)}const jt={key:0,class:"popupmodal"},kt={class:"popupmodal__window"};function xt(t,e,o,n,r,a){return(0,h.wg)(),(0,h.j4)(s.uT,{name:"fade"},{default:(0,h.w5)((()=>[t.isVisible?((0,h.wg)(),(0,h.iD)("div",jt,[(0,h._)("div",kt,[(0,h.WI)(t.$slots,"default")])])):(0,h.kq)("",!0)])),_:3})}var St={name:"PopupModal",data:()=>({isVisible:!1}),methods:{open(){this.isVisible=!0},close(){this.isVisible=!1}}};const Ct=(0,I.Z)(St,[["render",xt]]);var Jt=Ct,Pt={name:"ConfirmDialogue",components:{PopupModal:Jt},data:()=>({title:void 0,message:void 0,image:void 0,okButton:void 0,cancelButton:"Cancel",resolvePromise:void 0,rejectPromise:void 0}),methods:{show(t={}){return this.title=t.title,this.message=t.message,this.image=t.image,this.okButton=t.okButton,t.cancelButton&&(this.cancelButton=t.cancelButton),this.$refs.popup.open(),new Promise(((t,e)=>{this.resolvePromise=t,this.rejectPromise=e}))},_confirm(){this.$refs.popup.close(),this.resolvePromise(!0)},_cancel(){this.$refs.popup.close(),this.resolvePromise(!1)}}};const Dt=(0,I.Z)(Pt,[["render",yt]]);var Tt=Dt,Ot={name:"StatusItemComponent",components:{ConfirmDialogue:Tt},props:{job:{type:Object,required:!0,iteration_max:{type:String,required:!0},iteration_status:{type:String,required:!0},job_status:{type:String,required:!0},job_id:{type:String,required:!0},prompt:{type:String,required:!0}}},methods:{onClickSetSelected(t){this.$store.dispatch("getSelectedJob",t),this.buildModalDialogue()},async buildModalDialogue(){console.log("triggered");await this.$refs.confirmDialogue.show({title:this.getSelectedJob.prompt,image:this.getSelectedJob.img,message:"",okButton:"Löschen"})},getJobBorderClass(t){let e;switch(t){case"completed":e="border-success";break;case"processing":e="border-info";break;case"queued":e="border-warning";break;default:break}return e},getProgressbarPercent(t,e){return t/e*100}},computed:{getSelectedJob(){return this.$store.getters.getSelectedJob}}};const Et=(0,I.Z)(Ot,[["render",vt],["__scopeId","data-v-ec877518"]]);var Ht=Et,$t={name:"StatusListComponent",components:{StatusItem:Ht},data(){return{searchQuery:""}},methods:{getFoundJobs(t){return t.filter((t=>-1!=t.prompt.toLowerCase().indexOf(this.searchQuery.toLowerCase())))}},computed:{getJobs(){return this.$store.getters.getJobs},getFilteredJobs(){let t=this.getJobs;return t.sort((t=>"completed"==t.job_status)),this.getFoundJobs(t)}},watch:{getFilteredJobs:{handler(t){t&&t.forEach((t=>{"accepted"==t.job_status&&setTimeout((()=>{this.$store.dispatch("fetchJobs")}),1500)}))}}},async mounted(){this.$store.dispatch("fetchJobs")}};const Bt=(0,I.Z)($t,[["render",ct]]);var It=Bt,Wt={name:"HomeComponent",components:{Typing:L,Instructions:R,Prompt:st,StatusItemList:It},props:{msg:String}};const Lt=(0,I.Z)(Wt,[["render",D]]);var Nt=Lt;const Ft=[{path:"/",name:"Home",component:Nt}],Mt=(0,v.p7)({history:(0,v.PO)("/"),routes:Ft});var Zt=Mt;const qt={class:"appContainer"},zt={class:"row"},Ut={class:"col-12"};function Vt(t,e,o,s,n,r){const a=(0,h.up)("Navbar"),i=(0,h.up)("Home"),l=(0,h.up)("Footer");return(0,h.wg)(),(0,h.iD)("div",qt,[(0,h.Wm)(a),(0,h._)("div",zt,[(0,h._)("div",Ut,[(0,h.Wm)(i)])]),(0,h.Wm)(l)])}const Yt={class:"navbar navbar-expand-lg navbar-light fixed-top"},At={class:"container"},Qt=(0,h._)("a",{class:"navbar-brand",href:"#"},"Exia",-1),Kt=(0,h._)("button",{class:"navbar-toggler",type:"button","data-bs-toggle":"collapse","data-bs-target":"#navbarNav","aria-controls":"navbarNav","aria-expanded":"false","aria-label":"Toggle navigation",style:{width:"auto"}},[(0,h._)("span",{class:"navbar-toggler-icon"})],-1),Gt={class:"collapse navbar-collapse",id:"navbarNav"},Rt={class:"navbar-nav"},Xt={class:"nav-item"},te=(0,h.Uk)("Home"),ee={class:"nav-item"},oe=(0,h.Uk)("Settings"),se={class:"nav-item"},ne=(0,h.Uk)("About"),re={class:"nav-item"},ae=(0,h.Uk)("Login");function ie(t,e,o,s,n,r){const a=(0,h.up)("router-link");return(0,h.wg)(),(0,h.iD)("nav",Yt,[(0,h._)("div",At,[Qt,Kt,(0,h._)("div",Gt,[(0,h._)("ul",Rt,[(0,h._)("li",Xt,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start active"},{default:(0,h.w5)((()=>[te])),_:1})]),(0,h._)("div",ee,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[oe])),_:1})]),(0,h._)("div",se,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[ne])),_:1})]),(0,h._)("div",re,[(0,h.Wm)(a,{to:"/",class:"nav-link text-start"},{default:(0,h.w5)((()=>[ae])),_:1})])])])])])}var le={name:"NavbarComponent"};const ce=(0,I.Z)(le,[["render",ie]]);var ue=ce;const de={class:"footer text-muted mt-auto py-3 bg-light"},me={class:"text-center p-3"},pe=(0,h.Uk)(" Made with "),ge=(0,h._)("i",{class:"fa fa-heart","aria-hidden":"true"},null,-1);function be(t,e,o,s,n,r){return(0,h.wg)(),(0,h.iD)("footer",de,[(0,h._)("div",me,[pe,ge,(0,h.Uk)(" by Valar and Zendo in "+(0,T.zw)(r.getYear),1)])])}var ve={name:"FooterComponent",data(){return{}},computed:{getYear(){return(new Date).getFullYear()}}};const he=(0,I.Z)(ve,[["render",be]]);var fe=he,_e={name:"App",components:{Home:Nt,Navbar:ue,Footer:fe}};const we=(0,I.Z)(_e,[["render",Vt]]);var ye=we;let je=(0,s.ri)(ye);je.use(n),je.use(b),je.use(Zt),je.mount("#app")}},e={};function o(s){var n=e[s];if(void 0!==n)return n.exports;var r=e[s]={exports:{}};return t[s].call(r.exports,r,r.exports,o),r.exports}o.m=t,function(){var t=[];o.O=function(e,s,n,r){if(!s){var a=1/0;for(u=0;u<t.length;u++){s=t[u][0],n=t[u][1],r=t[u][2];for(var i=!0,l=0;l<s.length;l++)(!1&r||a>=r)&&Object.keys(o.O).every((function(t){return o.O[t](s[l])}))?s.splice(l--,1):(i=!1,r<a&&(a=r));if(i){t.splice(u--,1);var c=n();void 0!==c&&(e=c)}}return e}r=r||0;for(var u=t.length;u>0&&t[u-1][2]>r;u--)t[u]=t[u-1];t[u]=[s,n,r]}}(),function(){o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,{a:e}),e}}(),function(){o.d=function(t,e){for(var s in e)o.o(e,s)&&!o.o(t,s)&&Object.defineProperty(t,s,{enumerable:!0,get:e[s]})}}(),function(){o.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})}}(),function(){var t={143:0};o.O.j=function(e){return 0===t[e]};var e=function(e,s){var n,r,a=s[0],i=s[1],l=s[2],c=0;if(a.some((function(e){return 0!==t[e]}))){for(n in i)o.o(i,n)&&(o.m[n]=i[n]);if(l)var u=l(o)}for(e&&e(s);c<a.length;c++)r=a[c],o.o(t,r)&&t[r]&&t[r][0](),t[r]=0;return o.O(u)},s=self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[];s.forEach(e.bind(null,0)),s.push=e.bind(null,s.push.bind(s))}();var s=o.O(void 0,[998],(function(){return o(1255)}));s=o.O(s)})();
//# sourceMappingURL=app.d3d603ac.js.map