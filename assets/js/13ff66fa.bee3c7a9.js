"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[288],{5191:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>r,default:()=>u,frontMatter:()=>i,metadata:()=>a,toc:()=>d});var o=n(4848),s=n(8453);const i={sidebar_position:1},r="Relations",a={id:"relations",title:"Relations",description:"To enable read with relations included use annotations.Edge().EnableRead()",source:"@site/docs/relations.md",sourceDirName:".",slug:"/relations",permalink:"/mkgo/docs/relations",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Getting started",permalink:"/mkgo/docs/intro"},next:{title:"Functionality tables",permalink:"/mkgo/docs/functionality-list"}},c={},d=[];function l(e){const t={code:"code",h1:"h1",p:"p",pre:"pre",strong:"strong",...(0,s.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(t.h1,{id:"relations",children:"Relations"}),"\n",(0,o.jsxs)(t.p,{children:["To enable read with relations included use ",(0,o.jsx)(t.strong,{children:"annotations.Edge().EnableRead()"})]}),"\n",(0,o.jsx)(t.pre,{children:(0,o.jsx)(t.code,{className:"language-go",metastring:"title='your_model_schema.go' {5}",children:'// Edges of the Cat.\nfunc (Cat) Edges() []ent.Edge {\n\treturn []ent.Edge{\n\t\tedge.To("kittens", Kitten.Type).Annotations(\n\t\t\tannotations.Edge().EnableRead(),\n\t\t),\n\t}\n}\n'})})]})}function u(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(l,{...e})}):l(e)}},8453:(e,t,n)=>{n.d(t,{R:()=>r,x:()=>a});var o=n(6540);const s={},i=o.createContext(s);function r(e){const t=o.useContext(i);return o.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function a(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:r(e.components),o.createElement(i.Provider,{value:t},e.children)}}}]);