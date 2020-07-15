import { Component, OnInit, Input } from '@angular/core';
import { Thread } from '@models/thread.model';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import { ThreadComponent } from '../thread/thread.component';

@Component({
  selector: 'app-thread-card',
  templateUrl: './thread-card.component.html',
  styleUrls: ['./thread-card.component.scss']
})
export class ThreadCardComponent implements OnInit {
  @Input() thread: Thread;

  constructor(private dialog: MatDialog) { }

  ngOnInit(): void {
    console.log(this.thread);
  }

  public openThread = (thread: Thread): void => {
    const threadDialog = this.dialog.open(ThreadComponent);
    const instance = threadDialog.componentInstance;
    instance.thread = this.thread;
    instance.ngOnInit();
  }
}
