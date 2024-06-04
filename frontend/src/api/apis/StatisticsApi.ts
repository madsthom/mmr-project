/* tslint:disable */
/* eslint-disable */
/**
 * 
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ViewPlayerHistoryDetails,
} from '../models/index';
import {
    ViewPlayerHistoryDetailsFromJSON,
    ViewPlayerHistoryDetailsToJSON,
} from '../models/index';

export interface V1StatsPlayerHistoryUserIdGetRequest {
    userId: number;
    start?: string;
    end?: string;
}

/**
 * 
 */
export class StatisticsApi extends runtime.BaseAPI {

    /**
     * Get player history including MMR and date
     * Get player history
     */
    async v1StatsPlayerHistoryUserIdGetRaw(requestParameters: V1StatsPlayerHistoryUserIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<ViewPlayerHistoryDetails>>> {
        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling v1StatsPlayerHistoryUserIdGet().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['start'] != null) {
            queryParameters['start'] = requestParameters['start'];
        }

        if (requestParameters['end'] != null) {
            queryParameters['end'] = requestParameters['end'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/stats/player-history/{userId}`.replace(`{${"userId"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ViewPlayerHistoryDetailsFromJSON));
    }

    /**
     * Get player history including MMR and date
     * Get player history
     */
    async v1StatsPlayerHistoryUserIdGet(requestParameters: V1StatsPlayerHistoryUserIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<ViewPlayerHistoryDetails>> {
        const response = await this.v1StatsPlayerHistoryUserIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
