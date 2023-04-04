import { IonicModule } from '@ionic/angular';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { EditorPage } from './editor.page';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';

import {EditorPageRoutingModule} from './editor-routing.module'

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ExploreContainerComponentModule,
    EditorPageRoutingModule
  ],
  declarations: [EditorPage]
})
export class EditorPageModule {}