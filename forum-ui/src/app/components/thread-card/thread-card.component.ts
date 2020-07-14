import { Component, OnInit, Input } from '@angular/core';
import { Thread } from '@models/thread.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-thread-card',
  templateUrl: './thread-card.component.html',
  styleUrls: ['./thread-card.component.scss']
})
export class ThreadCardComponent implements OnInit {
  @Input() thread: Thread;

  constructor(private router: Router) { }

  ngOnInit(): void {
    console.log(this.thread);
  }

  public openThread = (): void => {
    console.log(this.thread.id);
    this.router.navigate(['/thread', { id: this.thread.id.toString() }]);
  }
}
