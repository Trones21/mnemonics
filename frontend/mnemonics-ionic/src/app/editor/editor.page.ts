import { Component, Injectable, Input ,OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MnemonicService } from '../mnemonic/mnemonic.service';
import { IonicModule } from '@ionic/angular';
import { CommonModule } from '@angular/common';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';




@Component({
  selector: 'app-editor',
  standalone: true,
  templateUrl: './editor.page.html',
  styleUrls: ['./editor.page.css'],
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ExploreContainerComponentModule,
  ],

})
export class EditorPage {
  
  form: FormGroup;

  constructor(private fb: FormBuilder, private mnemonicService: MnemonicService){
    this.form = this.fb.group({
      name: '',
      hint: '',
      answer: ''
    })
  }

  save(){

    //console.log(this.form.value)
    this.mnemonicService.createMnemonic(this.form.value)
    this.form.reset()
  }

  clearForm(){
    this.form.reset()

  }

}
