/*
 * Copyright © 2018. TIBCO Software Inc. [JGR]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
import {Injectable} from "@angular/core";
import {WiContrib, WiServiceHandlerContribution, AUTHENTICATION_TYPE} from "wi-studio/app/contrib/wi-contrib";
import {IConnectorContribution, IFieldDefinition, IActionResult, ActionResult} from "wi-studio/common/models/contrib";
import {Observable} from "rxjs/Observable";
import {IValidationResult, ValidationResult, ValidationError} from "wi-studio/common/models/validation";

@WiContrib({})
@Injectable()
export class TibcoSMTPConnectorContribution extends WiServiceHandlerContribution {
    constructor() {
        super();
    }

    value = (fieldName: string, context: IConnectorContribution): Observable<any> | any => {
        return null;
    }

    validate = (name: string, context: IConnectorContribution): Observable<IValidationResult> | IValidationResult => {
      if (name === "Connect") {
         let Server: IFieldDefinition;

         for (let configuration of context.settings) {
            if (configuration.name === "Server") {
                Server = configuration;
            }
    }

         if (Server.value) {
            // Enable Connect button
            return ValidationResult.newValidationResult().setReadOnly(false)
         } else {
            return ValidationResult.newValidationResult().setReadOnly(true)
         }
      }
       return null;
    }

    action = (actionName: string, context: IConnectorContribution): Observable<IActionResult> | IActionResult => {
       if (actionName == "Connect") {
          return Observable.create(observer => {
            let Server: IFieldDefinition;

            for (let configuration of context.settings) {
                if (configuration.name === "Server") {
                    Server = configuration;
                }
            }
			// Connect Test here. ... skipped for now!

			// Successfully connected. Lets save the configuration.	
            let actionResult = {
                context: context,
                authType: AUTHENTICATION_TYPE.BASIC,
                authData: {}
            };
            observer.next(ActionResult.newActionResult().setSuccess(true).setResult(actionResult));
        });
       }
       return null;
    }
}