import { Component, OnInit } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { UserService, User } from '../services/user.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { AddUserDialogComponent } from '../components/add-user-dialog/add-user-dialog.component';
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
    'id', 
    'user_name', 
    'first_name', 
    'last_name', 
    'email', 
    'user_status', 
    'department', 
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
      next: (users) => {
        this.users = users;
        this.isLoading = false;
      },
      error: (error) => {
        this.isLoading = false;
        this.showError('Failed to load users. The API server might be down or unreachable.');
        this.users = []; // Clear the users array when there's an error
      }
    });
  }

  private showError(message: string) {
    this.snackBar.open(message, 'Close', {
      duration: 5000,  // Show for 5 seconds
      horizontalPosition: 'center',
      verticalPosition: 'bottom',
      panelClass: ['error-snackbar']
    });
  }

  addUser() {
    const dialogRef = this.dialog.open(AddUserDialogComponent, {
      data: { user: null }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.userService.addUser(result).subscribe({
          next: (newUser) => {
            this.users = [...this.users, newUser];
            console.log('User added successfully');
          },
          error: (error) => {
            console.error('Error adding user:', error);
          }
        });
      }
    });
  }

  editUser(user: User) {
    const dialogRef = this.dialog.open(AddUserDialogComponent, {
      data: { user: { ...user } }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        const updatedUser = { ...result, id: user.id };
        this.userService.updateUser(updatedUser).subscribe({
          next: (response) => {
            const index = this.users.findIndex(u => u.id === user.id);
            if (index !== -1) {
              this.users[index] = response;
              this.users = [...this.users];
            }
            console.log('User updated successfully');
          },
          error: (error) => {
            console.error('Error updating user:', error);
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
        this.userService.deleteUser(user.id).subscribe({
          next: () => {
            this.users = this.users.filter(u => u.id !== user.id);
            console.log('User deleted successfully');
          },
          error: (error) => {
            console.error('Error deleting user:', error);
          }
        });
      }
    });
  }
}