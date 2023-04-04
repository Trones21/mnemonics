import { Component, Input ,OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';



class Mnemonic{
  name: string;
  hint: string;
  answer: string;
  constructor(
    name: string,
    hint: string,
    answer: string
  ){
    this.name = name;
    this.hint = hint;
    this.answer = answer;
  }

}

@Component({
  selector: 'app-editor',
  templateUrl: './editor.page.html',
  styleUrls: ['./editor.page.css'],
})
export class EditorPage {
  
  form: FormGroup;

  constructor(private fb: FormBuilder){
    this.form = this.fb.group({
      name: '',
      hint: '',
      answer: ''
    })
  }

  save(){
    console.log(this.form.value)
    
    //this.form.reset()
  }

  clearForm(){
    this.form.reset()

  }

}
