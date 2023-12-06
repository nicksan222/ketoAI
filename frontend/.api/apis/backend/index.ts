import type * as types from './types';
import type { ConfigOptions, FetchResponse } from 'api/dist/core'
import Oas from 'oas';
import APICore from 'api/dist/core';
import definition from './openapi.json';

class SDK {
  spec: Oas;
  core: APICore;

  constructor() {
    this.spec = Oas.init(definition);
    this.core = new APICore(this.spec, 'backend/1.0 (api/6.1.1)');
  }

  /**
   * Optionally configure various options that the SDK allows.
   *
   * @param config Object of supported SDK options and toggles.
   * @param config.timeout Override the default `fetch` request timeout of 30 seconds. This number
   * should be represented in milliseconds.
   */
  config(config: ConfigOptions) {
    this.core.setConfig(config);
  }

  /**
   * If the API you're using requires authentication you can supply the required credentials
   * through this method and the library will magically determine how they should be used
   * within your API request.
   *
   * With the exception of OpenID and MutualTLS, it supports all forms of authentication
   * supported by the OpenAPI specification.
   *
   * @example <caption>HTTP Basic auth</caption>
   * sdk.auth('username', 'password');
   *
   * @example <caption>Bearer tokens (HTTP or OAuth 2)</caption>
   * sdk.auth('myBearerToken');
   *
   * @example <caption>API Keys</caption>
   * sdk.auth('myApiKey');
   *
   * @see {@link https://spec.openapis.org/oas/v3.0.3#fixed-fields-22}
   * @see {@link https://spec.openapis.org/oas/v3.1.0#fixed-fields-22}
   * @param values Your auth credentials for the API; can specify up to two strings or numbers.
   */
  auth(...values: string[] | number[]) {
    this.core.setAuth(...values);
    return this;
  }

  /**
   * If the API you're using offers alternate server URLs, and server variables, you can tell
   * the SDK which one to use with this method. To use it you can supply either one of the
   * server URLs that are contained within the OpenAPI definition (along with any server
   * variables), or you can pass it a fully qualified URL to use (that may or may not exist
   * within the OpenAPI definition).
   *
   * @example <caption>Server URL with server variables</caption>
   * sdk.server('https://{region}.api.example.com/{basePath}', {
   *   name: 'eu',
   *   basePath: 'v14',
   * });
   *
   * @example <caption>Fully qualified server URL</caption>
   * sdk.server('https://eu.api.example.com/v14');
   *
   * @param url Server URL
   * @param variables An object of variables to replace into the server URL.
   */
  server(url: string, variables = {}) {
    this.core.setServer(url, variables);
  }

  /**
   * Retrieves a list of ingredients based on beginning and ending characters and a limit on
   * the number of results.
   *
   * @summary List ingredients
   * @throws FetchError<400, types.GetIngredientsResponse400> Bad Request - Invalid query parameters.
   * @throws FetchError<404, types.GetIngredientsResponse404> Not Found - Ingredients not found.
   */
  getIngredients(metadata?: types.GetIngredientsMetadataParam): Promise<FetchResponse<200, types.GetIngredientsResponse200>> {
    return this.core.fetch('/ingredients', 'get', metadata);
  }

  /**
   * Retrieves the ingredient by its ID.
   *
   * @summary Retrieve a single ingredient
   * @throws FetchError<400, types.GetIngredientsIngredientIdResponse400> Bad Request - Missing ingredient ID.
   * @throws FetchError<404, types.GetIngredientsIngredientIdResponse404> Not Found - Ingredient not found.
   */
  getIngredientsIngredient_id(metadata: types.GetIngredientsIngredientIdMetadataParam): Promise<FetchResponse<200, types.GetIngredientsIngredientIdResponse200>> {
    return this.core.fetch('/ingredients/{ingredient_id}', 'get', metadata);
  }

  /**
   * Retrieves the list of ingredient preferences for a user based on their ID.
   *
   * @summary Retrieve a user's ingredient preferences
   * @throws FetchError<401, types.GetIngredientsPreferencesResponse401> Unauthorized - User ID not found.
   * @throws FetchError<404, types.GetIngredientsPreferencesResponse404> Not Found - Ingredient preferences not found.
   */
  getIngredientsPreferences(): Promise<FetchResponse<200, types.GetIngredientsPreferencesResponse200>> {
    return this.core.fetch('/ingredients/preferences', 'get');
  }

  /**
   * Sets or updates the ingredient preferences for a user based on their ID.
   *
   * @summary Set a user's ingredient preferences
   * @throws FetchError<400, types.PostIngredientsPreferencesResponse400> Bad Request - Invalid request body or parameters.
   * @throws FetchError<401, types.PostIngredientsPreferencesResponse401> Unauthorized - User ID not found.
   * @throws FetchError<404, types.PostIngredientsPreferencesResponse404> Not Found - Unable to set or update preferences.
   */
  postIngredientsPreferences(body: types.PostIngredientsPreferencesBodyParam): Promise<FetchResponse<200, types.PostIngredientsPreferencesResponse200>> {
    return this.core.fetch('/ingredients/preferences', 'post', body);
  }

  /**
   * Deletes the preference of a specific ingredient for a user based on their ID and the
   * ingredient ID.
   *
   * @summary Delete a user's ingredient preference
   * @throws FetchError<400, types.DeleteIngredientsPreferencesIngredientIdResponse400> Bad Request - Missing ingredient ID.
   * @throws FetchError<401, types.DeleteIngredientsPreferencesIngredientIdResponse401> Unauthorized - User ID not found.
   * @throws FetchError<404, types.DeleteIngredientsPreferencesIngredientIdResponse404> Not Found - Ingredient or preference not found.
   */
  deleteIngredientsPreferencesIngredient_id(metadata: types.DeleteIngredientsPreferencesIngredientIdMetadataParam): Promise<FetchResponse<200, types.DeleteIngredientsPreferencesIngredientIdResponse200>> {
    return this.core.fetch('/ingredients/preferences/{ingredient_id}', 'delete', metadata);
  }
}

const createSDK = (() => { return new SDK(); })()
;

export default createSDK;

export type { DeleteIngredientsPreferencesIngredientIdMetadataParam, DeleteIngredientsPreferencesIngredientIdResponse200, DeleteIngredientsPreferencesIngredientIdResponse400, DeleteIngredientsPreferencesIngredientIdResponse401, DeleteIngredientsPreferencesIngredientIdResponse404, GetIngredientsIngredientIdMetadataParam, GetIngredientsIngredientIdResponse200, GetIngredientsIngredientIdResponse400, GetIngredientsIngredientIdResponse404, GetIngredientsMetadataParam, GetIngredientsPreferencesResponse200, GetIngredientsPreferencesResponse401, GetIngredientsPreferencesResponse404, GetIngredientsResponse200, GetIngredientsResponse400, GetIngredientsResponse404, PostIngredientsPreferencesBodyParam, PostIngredientsPreferencesResponse200, PostIngredientsPreferencesResponse400, PostIngredientsPreferencesResponse401, PostIngredientsPreferencesResponse404 } from './types';
