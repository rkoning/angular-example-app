import { Component, OnInit, Input } from '@angular/core';
import { Thread } from '@models/thread.model';
import { ThreadRoutesService } from '@routes/thread-routes.service';
import { ThreadService } from '@services/thread.service';
import { Comment } from '@models/comment.model';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.scss']
})
export class ThreadComponent implements OnInit {
  @Input() thread: Thread;

  public comments: Comment[];

  constructor(private threadService: ThreadService, public dialog: MatDialog) { }

  ngOnInit(): void {
    this.threadService.Comments.subscribe((comments) => {
      this.comments = comments;
    });
    this.threadService.getThreadComments(this.thread.id).subscribe((res) => {});
  }

  public deleteThread = (): void => {
    this.threadService.deleteThread(this.thread.id).subscribe((res) => {
      console.log(res);
    });
    this.dialog.closeAll();
  }
}
