package com.lomolo.copodapp.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.OutlinedIconButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.screens.UploadDisplayPictureDestination

@Composable
fun OnboardDpCard(
    modifier: Modifier = Modifier,
    onNext: (String) -> Unit,
    land: String,
    govtId: String,
) {
    OutlinedCard(
        modifier
            .fillMaxWidth()
            .height(360.dp)
    ) {
        Box(
            Modifier
                .fillMaxSize()
                .padding(12.dp)
        ) {
            Column(
                Modifier.align(Alignment.TopStart),
            ) {
                Icon(
                    painterResource(R.drawable.account),
                    modifier = Modifier.size(32.dp),
                    contentDescription = stringResource(R.string.display_picture),
                )
                Spacer(Modifier.size(20.dp))
                Text(
                    stringResource(R.string.display_picture),
                    style = MaterialTheme.typography.displaySmall,
                )
                Text(
                    stringResource(R.string.picture_of_you)
                )
            }
            Text(
                stringResource(R.string.upload_dp_copy_text),
                modifier = Modifier.align(Alignment.CenterStart),
            )
            OutlinedIconButton(
                onClick = { onNext("${UploadDisplayPictureDestination.route}/${false}") },
                modifier = Modifier
                    .align(Alignment.BottomEnd)
                    .size(60.dp),
                enabled = land.isNotEmpty() && govtId.isNotEmpty(),
            ) {
                Icon(
                    painterResource(R.drawable.doc_paper),
                    modifier = Modifier.size(32.dp),
                    contentDescription = stringResource(R.string.upload)
                )
            }
        }
    }
}