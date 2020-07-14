import { Component, OnInit, Input } from '@angular/core';
import { Thread } from '@models/thread.model';
import { ThreadRoutesService } from '@routes/thread-routes.service';
import { Router, ActivatedRoute } from '@angular/router';
import { ThreadService } from '@services/thread.service';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.scss']
})
export class ThreadComponent implements OnInit {
  @Input() thread: Thread;

  public comments: Comment[];
  public showSpinner = true;
  constructor(private route: ActivatedRoute, private threadService: ThreadService, private threadRoutesService: ThreadRoutesService) { }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.thread = this.threadService.getThreadById(id);
    console.log(this.thread);
    this.threadRoutesService.getComments(id).subscribe((res) => {

    });
    setTimeout(() => {
      this.showSpinner = false;
    }, 3000);
  }
}
