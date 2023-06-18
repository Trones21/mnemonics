import { Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss'],
})
export class AppComponent implements OnInit{
  constructor() {}

  ngOnInit(){
    localStorage.setItem('loggedInUserId', '6425fe608d3ae9661a039848')
  }
}
