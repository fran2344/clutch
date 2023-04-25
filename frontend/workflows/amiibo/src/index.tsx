import type { BaseWorkflowProps, WorkflowConfiguration } from "@clutch-sh/core";

import HelloWorld from "./hello-world";

export interface WorkflowProps extends BaseWorkflowProps {}

const register = (): WorkflowConfiguration => {
  return {
    developer: {
      name: "Name McName",
      contactUrl: "mailto:foo@foo-email.com",
    },
    path: "amiibo",
    group: "Amiibo",
    displayName: "Amiibo",
    routes: {
      landing: {
        path: "/lookup",
        description: "Lookup all Amiibo by name.",
        component: HelloWorld,
      },
    },
  };
};

export default register;
