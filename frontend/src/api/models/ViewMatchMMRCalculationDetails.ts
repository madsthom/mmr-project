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

import { mapValues } from '../runtime';
import type { ViewMatchMMRCalculationTeam } from './ViewMatchMMRCalculationTeam';
import {
    ViewMatchMMRCalculationTeamFromJSON,
    ViewMatchMMRCalculationTeamFromJSONTyped,
    ViewMatchMMRCalculationTeamToJSON,
} from './ViewMatchMMRCalculationTeam';

/**
 * 
 * @export
 * @interface ViewMatchMMRCalculationDetails
 */
export interface ViewMatchMMRCalculationDetails {
    /**
     * 
     * @type {ViewMatchMMRCalculationTeam}
     * @memberof ViewMatchMMRCalculationDetails
     */
    team1: ViewMatchMMRCalculationTeam;
    /**
     * 
     * @type {ViewMatchMMRCalculationTeam}
     * @memberof ViewMatchMMRCalculationDetails
     */
    team2: ViewMatchMMRCalculationTeam;
}

/**
 * Check if a given object implements the ViewMatchMMRCalculationDetails interface.
 */
export function instanceOfViewMatchMMRCalculationDetails(value: object): boolean {
    if (!('team1' in value)) return false;
    if (!('team2' in value)) return false;
    return true;
}

export function ViewMatchMMRCalculationDetailsFromJSON(json: any): ViewMatchMMRCalculationDetails {
    return ViewMatchMMRCalculationDetailsFromJSONTyped(json, false);
}

export function ViewMatchMMRCalculationDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ViewMatchMMRCalculationDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'team1': ViewMatchMMRCalculationTeamFromJSON(json['team1']),
        'team2': ViewMatchMMRCalculationTeamFromJSON(json['team2']),
    };
}

export function ViewMatchMMRCalculationDetailsToJSON(value?: ViewMatchMMRCalculationDetails | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'team1': ViewMatchMMRCalculationTeamToJSON(value['team1']),
        'team2': ViewMatchMMRCalculationTeamToJSON(value['team2']),
    };
}
