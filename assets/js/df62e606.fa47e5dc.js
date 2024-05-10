"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[8089],{1914:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>c,contentTitle:()=>o,default:()=>u,frontMatter:()=>a,metadata:()=>s,toc:()=>d});var r=n(85893),t=n(11151);const a={title:"Upgrading IBC Chains Overview",sidebar_label:"Overview",sidebar_position:0,slug:"/ibc/upgrades/intro"},o=void 0,s={id:"ibc/upgrades/intro",title:"Upgrading IBC Chains Overview",description:"Upgrading IBC Chains Overview",source:"@site/versioned_docs/version-v7.4.x/01-ibc/05-upgrades/00-intro.md",sourceDirName:"01-ibc/05-upgrades",slug:"/ibc/upgrades/intro",permalink:"/v7/ibc/upgrades/intro",draft:!1,unlisted:!1,tags:[],version:"v7.4.x",sidebarPosition:0,frontMatter:{title:"Upgrading IBC Chains Overview",sidebar_label:"Overview",sidebar_position:0,slug:"/ibc/upgrades/intro"},sidebar:"defaultSidebar",previous:{title:"Integrating IBC middleware into a chain",permalink:"/v7/ibc/middleware/integration"},next:{title:"How to Upgrade IBC Chains and their Clients",permalink:"/v7/ibc/upgrades/quick-guide"}},c={},d=[{value:"Upgrading IBC Chains Overview",id:"upgrading-ibc-chains-overview",level:3}];function l(e){const i={a:"a",h3:"h3",li:"li",ol:"ol",p:"p",...(0,t.a)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(i.h3,{id:"upgrading-ibc-chains-overview",children:"Upgrading IBC Chains Overview"}),"\n",(0,r.jsx)(i.p,{children:"This directory contains information on how to upgrade an IBC chain without breaking counterparty clients and connections."}),"\n",(0,r.jsx)(i.p,{children:"IBC-connnected chains must be able to upgrade without breaking connections to other chains. Otherwise there would be a massive disincentive towards upgrading and disrupting high-value IBC connections, thus preventing chains in the IBC ecosystem from evolving and improving. Many chain upgrades may be irrelevant to IBC, however some upgrades could potentially break counterparty clients if not handled correctly. Thus, any IBC chain that wishes to perform a IBC-client-breaking upgrade must perform an IBC upgrade in order to allow counterparty clients to securely upgrade to the new light client."}),"\n",(0,r.jsxs)(i.ol,{children:["\n",(0,r.jsxs)(i.li,{children:["The ",(0,r.jsx)(i.a,{href:"/v7/ibc/upgrades/quick-guide",children:"quick-guide"})," describes how IBC-connected chains can perform client-breaking upgrades and how relayers can securely upgrade counterparty clients using the SDK."]}),"\n",(0,r.jsxs)(i.li,{children:["The ",(0,r.jsx)(i.a,{href:"/v7/ibc/upgrades/developer-guide",children:"developer-guide"})," is a guide for developers intending to develop IBC client implementations with upgrade functionality."]}),"\n"]})]})}function u(e={}){const{wrapper:i}={...(0,t.a)(),...e.components};return i?(0,r.jsx)(i,{...e,children:(0,r.jsx)(l,{...e})}):l(e)}},11151:(e,i,n)=>{n.d(i,{Z:()=>s,a:()=>o});var r=n(67294);const t={},a=r.createContext(t);function o(e){const i=r.useContext(a);return r.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function s(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:o(e.components),r.createElement(a.Provider,{value:i},e.children)}}}]);