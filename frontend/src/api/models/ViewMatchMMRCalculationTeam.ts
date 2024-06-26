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
/**
 * 
 * @export
 * @interface ViewMatchMMRCalculationTeam
 */
export interface ViewMatchMMRCalculationTeam {
    /**
     * 
     * @type {number}
     * @memberof ViewMatchMMRCalculationTeam
     */
    player1MMRDelta: number;
    /**
     * 
     * @type {number}
     * @memberof ViewMatchMMRCalculationTeam
     */
    player2MMRDelta: number;
}

/**
 * Check if a given object implements the ViewMatchMMRCalculationTeam interface.
 */
export function instanceOfViewMatchMMRCalculationTeam(value: object): boolean {
    if (!('player1MMRDelta' in value)) return false;
    if (!('player2MMRDelta' in value)) return false;
    return true;
}

export function ViewMatchMMRCalculationTeamFromJSON(json: any): ViewMatchMMRCalculationTeam {
    return ViewMatchMMRCalculationTeamFromJSONTyped(json, false);
}

export function ViewMatchMMRCalculationTeamFromJSONTyped(json: any, ignoreDiscriminator: boolean): ViewMatchMMRCalculationTeam {
    if (json == null) {
        return json;
    }
    return {
        
        'player1MMRDelta': json['player1MMRDelta'],
        'player2MMRDelta': json['player2MMRDelta'],
    };
}

export function ViewMatchMMRCalculationTeamToJSON(value?: ViewMatchMMRCalculationTeam | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'player1MMRDelta': value['player1MMRDelta'],
        'player2MMRDelta': value['player2MMRDelta'],
    };
}

