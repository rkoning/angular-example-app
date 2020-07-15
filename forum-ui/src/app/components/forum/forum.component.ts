import { Component, OnInit } from '@angular/core';
import { ThreadService } from '@services/thread.service';
import { Thread } from '@models/thread.model';
import { MatDialog } from '@angular/material/dialog';
import { ThreadFormComponent } from '../thread-form/thread-form.component';
import { ThreadRoutesService } from '@routes/thread-routes.service';

@Component({
  selector: 'app-forum',
  templateUrl: './forum.component.html',
  styleUrls: ['./forum.component.scss']
})
export class ForumComponent implements OnInit {

  public threads: Thread[] = [];

  constructor(private threadService: ThreadService, public threadRoutes: ThreadRoutesService, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.threadService.Threads.subscribe((threads) => {
      this.threads = threads;
    });

    this.threadRoutes.list().subscribe((res) => {
      this.threadService.setThreads(res.data);
    });
  }

  public openAddThreadDialog = (): void => {
    const threadDialog = this.dialog.open(ThreadFormComponent);
  }

}
