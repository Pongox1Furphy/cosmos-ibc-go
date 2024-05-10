"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[3085],{16144:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>h,frontMatter:()=>a,metadata:()=>r,toc:()=>d});var i=n(85893),s=n(11151);const a={title:"Handling Genesis",sidebar_label:"Handling Genesis",sidebar_position:8,slug:"/ibc/light-clients/genesis"},o="Genesis metadata",r={id:"light-clients/developer-guide/genesis",title:"Handling Genesis",description:"Learn how to implement the ExportMetadata interface",source:"@site/versioned_docs/version-v8.2.x/03-light-clients/01-developer-guide/08-genesis.md",sourceDirName:"03-light-clients/01-developer-guide",slug:"/ibc/light-clients/genesis",permalink:"/v8/ibc/light-clients/genesis",draft:!1,unlisted:!1,tags:[],version:"v8.2.x",sidebarPosition:8,frontMatter:{title:"Handling Genesis",sidebar_label:"Handling Genesis",sidebar_position:8,slug:"/ibc/light-clients/genesis"},sidebar:"defaultSidebar",previous:{title:"Handling Proposals",permalink:"/v8/ibc/light-clients/proposals"},next:{title:"Setup",permalink:"/v8/ibc/light-clients/setup"}},l={},d=[{value:"Pre-requisite readings",id:"pre-requisite-readings",level:2}];function c(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",li:"li",p:"p",pre:"pre",ul:"ul",...(0,s.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(t.h1,{id:"genesis-metadata",children:"Genesis metadata"}),"\n",(0,i.jsx)(t.admonition,{title:"Synopsis",type:"note",children:(0,i.jsxs)(t.p,{children:["Learn how to implement the ",(0,i.jsx)(t.code,{children:"ExportMetadata"})," interface"]})}),"\n",(0,i.jsxs)(t.admonition,{type:"note",children:[(0,i.jsx)(t.h2,{id:"pre-requisite-readings",children:"Pre-requisite readings"}),(0,i.jsxs)(t.ul,{children:["\n",(0,i.jsx)(t.li,{children:(0,i.jsx)(t.a,{href:"https://docs.cosmos.network/v0.47/building-modules/genesis",children:"Cosmos SDK module genesis"})}),"\n"]})]}),"\n",(0,i.jsxs)(t.p,{children:[(0,i.jsx)(t.code,{children:"ClientState"})," instances are provided their own isolated and namespaced client store upon initialisation. ",(0,i.jsx)(t.code,{children:"ClientState"})," implementations may choose to store any amount of arbitrary metadata in order to verify counterparty consensus state and perform light client updates correctly."]}),"\n",(0,i.jsxs)(t.p,{children:["The ",(0,i.jsx)(t.code,{children:"ExportMetadata"})," method of the ",(0,i.jsxs)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/core/exported/client.go#L47",children:[(0,i.jsx)(t.code,{children:"ClientState"})," interface"]})," provides light client modules with the ability to persist metadata in genesis exports."]}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{className:"language-go",children:"ExportMetadata(clientStore sdk.KVStore) []GenesisMetadata\n"})}),"\n",(0,i.jsxs)(t.p,{children:[(0,i.jsx)(t.code,{children:"ExportMetadata"})," is provided the client store and returns an array of ",(0,i.jsx)(t.code,{children:"GenesisMetadata"}),". For maximum flexibility, ",(0,i.jsx)(t.code,{children:"GenesisMetadata"})," is defined as a simple interface containing two distinct ",(0,i.jsx)(t.code,{children:"Key"})," and ",(0,i.jsx)(t.code,{children:"Value"})," accessor methods."]}),"\n",(0,i.jsx)(t.pre,{children:(0,i.jsx)(t.code,{className:"language-go",children:"type GenesisMetadata interface {\n  // return store key that contains metadata without clientID-prefix\n  GetKey() []byte\n  // returns metadata value\n  GetValue() []byte\n}\n"})}),"\n",(0,i.jsxs)(t.p,{children:["This allows ",(0,i.jsx)(t.code,{children:"ClientState"})," instances to retrieve and export any number of key-value pairs which are maintained within the store in their raw ",(0,i.jsx)(t.code,{children:"[]byte"})," form."]}),"\n",(0,i.jsxs)(t.p,{children:["When a chain is started with a ",(0,i.jsx)(t.code,{children:"genesis.json"})," file which contains ",(0,i.jsx)(t.code,{children:"ClientState"})," metadata (for example, when performing manual upgrades using an exported ",(0,i.jsx)(t.code,{children:"genesis.json"}),") the ",(0,i.jsx)(t.code,{children:"02-client"})," submodule of core IBC will handle setting the key-value pairs within their respective client stores. ",(0,i.jsxs)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/core/02-client/genesis.go#L18-L22",children:["See ",(0,i.jsx)(t.code,{children:"02-client"})," ",(0,i.jsx)(t.code,{children:"InitGenesis"})]}),"."]}),"\n",(0,i.jsxs)(t.p,{children:["Please refer to the ",(0,i.jsx)(t.a,{href:"https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/light-clients/07-tendermint/genesis.go#L12",children:"Tendermint light client implementation"})," for an example."]})]})}function h(e={}){const{wrapper:t}={...(0,s.a)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(c,{...e})}):c(e)}},11151:(e,t,n)=>{n.d(t,{Z:()=>r,a:()=>o});var i=n(67294);const s={},a=i.createContext(s);function o(e){const t=i.useContext(a);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function r(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:o(e.components),i.createElement(a.Provider,{value:t},e.children)}}}]);