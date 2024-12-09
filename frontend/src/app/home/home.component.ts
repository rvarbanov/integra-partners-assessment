import { Component, OnInit } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { UserService, User } from '../services/user.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { UserDialogComponent } from '../components/user-dialog/user-dialog.component';
import { ConfirmDialogComponent } from '../components/confirm-dialog/confirm-dialog.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    MatTableModule, 
    MatButtonModule, 
    MatIconModule, 
    CommonModule,
    MatProgressSpinnerModule,
    MatSnackBarModule
  ],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  displayedColumns: string[] = [
    'user_id', 
    'user_name', 
    'first_name', 
    'last_name', 
    'email', 
    'user_status', 
    'department',
    'created_at',
    'updated_at',
    'actions'
  ];
  users: User[] = [];
  isLoading = false;

  constructor(
    private userService: UserService,
    private dialog: MatDialog,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit() {
    this.loadUsers();
  }

  loadUsers() {
    this.isLoading = true;
    this.userService.getUsers().subscribe({
      next: (data) => {
        this.users = data.sort((a, b) => a.user_id - b.user_id);
        this.isLoading = false;
      },
      error: (error) => {
        this.showError('Failed to load users. The API server might be down or unreachable.');
        console.error('error loading users:', error);
        this.isLoading = false;
        this.users = [];
      }
    });
  }

  private showError(message: string) {
    this.snackBar.open(message, 'Close', {
      duration: 5000,
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['error-snackbar']
    });
  }

  addUser() {
    const dialogRef = this.dialog.open(UserDialogComponent, {
      data: { user: null }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.addUser(result).subscribe({
          next: (response) => {
            this.snackBar.open('User added successfully', 'Close', {
              duration: 3000,
              horizontalPosition: 'center',
              verticalPosition: 'bottom'
            });
            this.loadUsers();
          },
          error: (error) => {
            console.error('Error adding user:', error);
            this.showError('Failed to add user: ' + (error.error?.error || 'Unknown error occurred'));
          }
        });
      }
    });
  }

  editUser(user: User) {
    const dialogRef = this.dialog.open(UserDialogComponent, {
      data: { user: { ...user } }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const updatedUser = { ...result, user_id: user.user_id };
        this.userService.updateUser(updatedUser).subscribe({
          next: (response) => {
            this.snackBar.open('User updated successfully', 'Close', {
              duration: 3000,
              horizontalPosition: 'center',
              verticalPosition: 'bottom'
            });
            this.loadUsers();
          },
          error: (error) => {
            console.error('Error updating user:', error);
            this.showError('Failed to update user: ' + (error.error?.error || 'Unknown error occurred'));
          }
        });
      }
    });
  }

  deleteUser(user: User) {
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      data: {
        title: 'Confirm Delete',
        message: `Are you sure you want to delete user ${user.first_name} ${user.last_name}?`
      }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.deleteUser(user.user_id).subscribe({
          next: () => {
            this.snackBar.open('User deleted successfully', 'Close', {
              duration: 3000,
              horizontalPosition: 'center',
              verticalPosition: 'bottom'
            });
            this.loadUsers();
          },
          error: (error) => {
            console.error('Error deleting user:', error);
            this.showError('Failed to delete user: ' + (error.error?.error || 'Unknown error occurred'));
          }
        });
      }
    });
  }
}