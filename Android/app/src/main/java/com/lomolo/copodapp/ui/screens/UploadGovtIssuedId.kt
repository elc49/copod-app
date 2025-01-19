package com.lomolo.copodapp.ui.screens

import android.provider.OpenableColumns
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowForward
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.state.viewmodels.UploadingDoc
import com.lomolo.copodapp.ui.common.UploadDocument
import com.lomolo.copodapp.ui.navigation.Navigation
import kotlinx.coroutines.launch

object UploadGovtIssuedIdScreenDestination : Navigation {
    override val title = null
    override val route = "register-govt-id"
}

@Composable
fun UploadGovtIssuedId(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    onNext: () -> Unit,
    viewModel: OnboardingViewModel,
) {
    val image by viewModel.supportingDoc.collectAsState()
    val idDoc = when (viewModel.uploadingGovtId) {
        UploadingDoc.Loading -> {
            R.drawable.loading_img
        }

        UploadingDoc.Success -> {
            if (image.isEmpty()) {
                R.drawable.upload
            } else {
                image
            }
        }

        else -> {
            R.drawable.upload
        }
    }
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val pickGovtIdMedia = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia()
    ) {
        if (it != null) {
            val stream = context.contentResolver.openInputStream(it)
            var fileName = "Unknown"
            context.contentResolver.query(it, null, null, null, null)?.use { cursor ->
                if (cursor.moveToFirst()) {
                    val nameIndex = cursor.getColumnIndex(OpenableColumns.DISPLAY_NAME)
                    if (nameIndex >= 0) {
                        fileName = cursor.getString(nameIndex) ?: "${System.currentTimeMillis()}"
                    }
                }
            }
            if (stream != null) {
                viewModel.uploadGovtIssuedId(fileName, stream)
            }
        }
    }

    UploadDocument(
        modifier = modifier,
        titleSize = "large",
        title = @Composable {
            Column {
                Text(stringResource(R.string.verify_your_id))
            }
        },
        newUpload = image.isEmpty(),
        image = idDoc,
        onNext = {
            if (image.isNotEmpty()) {
                onNext()
            }
        },
        onGoBack = onGoBack,
        onSelectImage = {
            if (viewModel.uploadingGovtId !is UploadingDoc.Loading) {
                scope.launch {
                    pickGovtIdMedia.launch(
                        PickVisualMediaRequest(
                            ActivityResultContracts.PickVisualMedia.ImageOnly,
                        )
                    )
                }
            }
        },
        savingDoc = viewModel.uploadingGovtId is UploadingDoc.Loading,
        buttonText = @Composable {
            Text(
                stringResource(R.string.proceed),
                style = MaterialTheme.typography.titleMedium,
            )
            Icon(
                Icons.AutoMirrored.TwoTone.ArrowForward,
                contentDescription = stringResource(R.string.proceed),
            )
        }
    )
}