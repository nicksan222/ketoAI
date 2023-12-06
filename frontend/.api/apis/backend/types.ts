import type { FromSchema } from 'json-schema-to-ts';
import * as schemas from './schemas';

export type GetIngredientsIngredientIdMetadataParam = FromSchema<typeof schemas.GetIngredientsIngredientId.metadata>;
export type GetIngredientsIngredientIdResponse200 = FromSchema<typeof schemas.GetIngredientsIngredientId.response['200']>;
export type GetIngredientsIngredientIdResponse400 = FromSchema<typeof schemas.GetIngredientsIngredientId.response['400']>;
export type GetIngredientsIngredientIdResponse404 = FromSchema<typeof schemas.GetIngredientsIngredientId.response['404']>;
