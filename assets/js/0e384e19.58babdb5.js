"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[976],{1512:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>r,default:()=>u,frontMatter:()=>o,metadata:()=>s,toc:()=>d});var i=t(4848),a=t(8453);const o={sidebar_position:0},r="Getting started",s={id:"intro",title:"Getting started",description:"MKGO is a tool aimed to reduce amount of work, required to start and maintain a go project.",source:"@site/docs/intro.md",sourceDirName:".",slug:"/intro",permalink:"/docs/intro",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:0,frontMatter:{sidebar_position:0},sidebar:"tutorialSidebar",next:{title:"Relations",permalink:"/docs/relations"}},l={},d=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Installation",id:"installation",level:2},{value:"Initialize project",id:"initialize-project",level:2},{value:"Create model",id:"create-model",level:2},{value:"Generate crud handlers",id:"generate-crud-handlers",level:2},{value:"Generate crud operations by default and if you want to disable some for your entity - do it with annotations.",id:"generate-crud-operations-by-default-and-if-you-want-to-disable-some-for-your-entity---do-it-with-annotations",level:3},{value:"Do not generate crud operations by default and if you want to enable some for your entity - do it with annotations.",id:"do-not-generate-crud-operations-by-default-and-if-you-want-to-enable-some-for-your-entity---do-it-with-annotations",level:3},{value:"Generate app",id:"generate-app",level:2},{value:"Run",id:"run",level:2}];function c(e){const n={a:"a",admonition:"admonition",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",li:"li",ol:"ol",p:"p",pre:"pre",...(0,a.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"getting-started",children:"Getting started"}),"\n",(0,i.jsx)(n.p,{children:"MKGO is a tool aimed to reduce amount of work, required to start and maintain a go project.\nIt provides commands to initialize project, add models(with ent), auto-generate migrations(with atlas),\nauto-generate crud operations with openapi schema."}),"\n",(0,i.jsxs)(n.ol,{children:["\n",(0,i.jsx)(n.li,{children:"install"}),"\n",(0,i.jsx)(n.li,{children:"initialize"}),"\n",(0,i.jsx)(n.li,{children:"describe your ent schema"}),"\n",(0,i.jsx)(n.li,{children:"generate"}),"\n",(0,i.jsx)(n.li,{children:"enjoy"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,i.jsx)(n.p,{children:"mkgo has dependencies:\natlas, ent, ogen, goimports"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"curl -sSf https://atlasgo.sh | ATLAS_VERSION=v0.19.3-cfa638c-canary sh\ngo install golang.org/x/tools/cmd/goimports@latest\n"})}),"\n",(0,i.jsx)(n.h2,{id:"installation",children:"Installation"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"go install github.com/wheissd/gomk/cmd/mkgo@latest\n"})}),"\n",(0,i.jsx)(n.admonition,{title:"TIP",type:"tip",children:(0,i.jsx)(n.p,{children:"Ensure your PATH contains go bin dir"})}),"\n",(0,i.jsx)(n.h2,{id:"initialize-project",children:"Initialize project"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"mkgo init hello_mkgo\n"})}),"\n",(0,i.jsx)(n.p,{children:"This creates basic project structure inside current folder"}),"\n",(0,i.jsx)(n.p,{children:"api\nent\nopenapi\ngen_config.yaml\ngenerate.go"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"project_root\n\u251c\u2500\u2500 internal\n\u2502   \u2514\u2500\u2500 api\n\u2502       \u251c\u2500\u2500 cmd\n\u2502       \u2502   \u2514\u2500\u2500 apigen\n\u2502       \u2502       \u251c\u2500\u2500 main.go\n\u2502       \u2502       \u2514\u2500\u2500 pre_gen.go\n\u2502       \u2514\u2500\u2500 gen\n\u2502           \u251c\u2500\u2500 cmd\n\u2502           \u2502   \u2514\u2500\u2500 main.go\n\u2502           \u251c\u2500\u2500 schema\n\u2502           \u2514\u2500\u2500 generate.go\n\u251c\u2500\u2500 ent\n\u251c\u2500\u2500 openapi\n\u251c\u2500\u2500 gen_config.yaml\n\u2514\u2500\u2500 generate.go\n"})}),"\n",(0,i.jsx)(n.h2,{id:"create-model",children:"Create model"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"mkgo model Example\n"})}),"\n",(0,i.jsxs)(n.p,{children:["Creates new ent model. For ent docs visit ",(0,i.jsx)(n.a,{href:"https://entgo.io/",children:"https://entgo.io/"})]}),"\n",(0,i.jsx)(n.h2,{id:"generate-crud-handlers",children:"Generate crud handlers"}),"\n",(0,i.jsx)(n.p,{children:"After you described your entities, you can auto-generate crud operations for them.\nThere are two options:"}),"\n",(0,i.jsxs)(n.ol,{children:["\n",(0,i.jsx)(n.li,{children:"Enable operations by default and disable them where you want."}),"\n",(0,i.jsx)(n.li,{children:"Disable operations by default and enable them where you want."}),"\n"]}),"\n",(0,i.jsx)(n.admonition,{title:"TIP",type:"tip",children:(0,i.jsx)(n.p,{children:"By default cruds are not generated and fields will not be exposed to generated cruds"})}),"\n",(0,i.jsx)(n.h3,{id:"generate-crud-operations-by-default-and-if-you-want-to-disable-some-for-your-entity---do-it-with-annotations",children:"Generate crud operations by default and if you want to disable some for your entity - do it with annotations."}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-yaml",metastring:"title='mkgo_config.yaml'",children:"    - rest-client:\n        OutputPath: api\n        OpenApiPath: openapi/rest_client.json\n        Mode: rest_client\n        Title: Your API title\n        Servers:\n          - http://localhost:9000\n        EnableDefaultReadOne: true\n        EnableDefaultReadMany: true\n        EnableDefaultCreate: true\n        EnableDefaultUpdate: true\n        EnableDefaultDelete: true\n        FieldsPublicByDefault: true\n"})}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-go",metastring:"title='your_model_schema.go' {13-28}",children:'package schema\n\nimport (\n\t"entgo.io/ent"\n\t"entgo.io/ent/schema/field"\n\t"github.com/wheissd/gomk/annotations"\n)\n\ntype Cat struct {\n\tent.Schema\n}\n\nfunc (Cat) Annotations() []schema.Annotation {\n\treturn []schema.Annotation{\n\t\tannotations.Entity().DisableDelete(annotations.Modes{"api"}),\n\t}\n}\n\nfunc (Cat) Fields() []ent.Field {\n\treturn []ent.Field{\n\t\t// this field will be expose in api\n\t\tfield.String("name"),\n                // this field wont be exposed in api\n\t\tfield.Int64("speed").Annotations(\n\t\t\tannotations.Field().SetPrivate(),\n\t\t),\n\t}\n}\n\n'})}),"\n",(0,i.jsx)(n.h3,{id:"do-not-generate-crud-operations-by-default-and-if-you-want-to-enable-some-for-your-entity---do-it-with-annotations",children:"Do not generate crud operations by default and if you want to enable some for your entity - do it with annotations."}),"\n",(0,i.jsx)(n.admonition,{title:"TIP",type:"tip",children:(0,i.jsx)(n.p,{children:"You can leave config options empty - false is default value"})}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-yaml",metastring:"title='mkgo_config.yaml'",children:"    - rest-client:\n        OutputPath: api\n        OpenApiPath: openapi/rest_client.json\n        Mode: rest_client\n        Title: Your API title\n        Servers:\n          - http://localhost:9000\n"})}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-go",metastring:"title='your_model_schema.go' {13-28}",children:'package schema\n\nimport (\n\t"entgo.io/ent"\n\t"entgo.io/ent/schema/field"\n\t"github.com/wheissd/gomk/annotations"\n)\n\ntype Cat struct {\n\tent.Schema\n}\n\nfunc (Cat) Annotations() []schema.Annotation {\n\treturn []schema.Annotation{\n\t\tannotations.Entity().EnableCreate(annotations.Modes{"api"}),\n\t}\n}\n\nfunc (Cat) Fields() []ent.Field {\n\treturn []ent.Field{\n\t\tfield.String("name").Annotations(\n\t\t\tannotations.Field().SetPublic(),\n\t\t),\n\t\tfield.Int64("speed").Annotations(\n\t\t\tannotations.Field().SetPublic(),\n\t\t),\n\t}\n}\n\n'})}),"\n",(0,i.jsx)(n.h2,{id:"generate-app",children:"Generate app"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"mkgo generate\n"})}),"\n",(0,i.jsx)(n.h2,{id:"run",children:"Run"}),"\n",(0,i.jsx)(n.admonition,{title:"Tip",type:"tip",children:(0,i.jsxs)(n.p,{children:["By default you'll have ",(0,i.jsx)(n.em,{children:"rest_client"})," and ",(0,i.jsx)(n.em,{children:"grpc_admin"})," as [your_generated_api_binary]"]})}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-console",children:"go run cmd/[your_generated_api_binary]\n"})}),"\n",(0,i.jsx)(n.p,{children:"Voila, now you can test and deploy your api"})]})}function u(e={}){const{wrapper:n}={...(0,a.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(c,{...e})}):c(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>r,x:()=>s});var i=t(6540);const a={},o=i.createContext(a);function r(e){const n=i.useContext(o);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function s(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:r(e.components),i.createElement(o.Provider,{value:n},e.children)}}}]);