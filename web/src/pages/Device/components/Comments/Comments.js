import { useState } from 'react';
import FormService from '@services/form.service';
import Container from '@components/Container/Container';
import Button from '@components/Button/Button';
import Input from '@components/Input/Input';
import CommentDisplay from './CommentDisplay';

export default function Comments({ data }) {
	const [comment, setComment] = useState('');

	const handleAddComm = async (e) => {
		e.preventDefault();
		FormService.addComment(data.Miner.IPAddress, comment).then(
			(response) => {
				console.log('Add comment ok');
				setComment('');
			},
			(error) => {
				console.log('Add comment', error);
			}
		);
	};

	const handleDeleteComment = async (
		event,
		username,
		content,
		creationDate
	) => {
		event.preventDefault();
		FormService.deleteComment(username, content, creationDate).then(
			(response) => {
				console.log('Delete comment ok');
			},
			(error) => {
				console.log('Delete comment ', error);
			}
		);
	};

	return (
		<div style={{ marginTop: '20px' }}>
			<Container>
				<div
					className="grid-50-50"
					style={{ columnGap: '20px', paddingInline: '20px' }}
				>
					<div>
						<div style={{ marginBottom: '20px' }}>Comments section</div>
						{data.Comments &&
							data.Comments.map((data, index) => (
								<div id={index}>
									<CommentDisplay
										username={data.Username}
										date={data.CreationDate}
										content={data.Content}
									/>
									<Button
										value="delete"
										float="right"
										onClick={(event) =>
											handleDeleteComment(
												event,
												data.Username,
												data.Content,
												data.CreationDate
											)
										}
									/>
								</div>
							))}
					</div>
					<div>
						<form onSubmit={handleAddComm} className="grid-85-15">
							<Input
								type="text"
								value={comment}
								setValue={setComment}
								width="fluid"
								placeholder="Type your comment here"
							/>
							<Button value="add comment" type="submit" />
						</form>
					</div>
				</div>
			</Container>
		</div>
	);
}
