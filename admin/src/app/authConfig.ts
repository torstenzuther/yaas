import { AuthConfig } from 'angular-oauth2-oidc';

export const authCodeFlowConfig: AuthConfig = {
  issuer: 'http://localhost',
  redirectUri: window.location.origin + '/index.html',
  clientId: 'torsten',
  responseType: 'code',
  scope: 'openid profile email offline_access',
  showDebugInformation: true,
  requireHttps: false,
};

