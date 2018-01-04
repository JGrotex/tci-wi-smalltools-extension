/*
 * Copyright © 2018. TIBCO Software Inc. [JGR]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
import { HttpModule } from "@angular/http";
import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import {TibcoSMTPConnectorContribution} from "./connector";
import {WiServiceContribution} from "wi-studio/app/contrib/wi-contrib";

@NgModule({
  imports: [
  	CommonModule,
  	HttpModule,
  ],
  providers: [
    {
       provide: WiServiceContribution,
       useClass: TibcoSMTPConnectorContribution
     }
  ]
})

export default class TibcoSMTPConnectorModule {

}